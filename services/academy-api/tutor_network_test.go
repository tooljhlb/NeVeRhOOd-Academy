package main

import (
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// Exercises real HTTP sockets on both sides of the Academy gateway.
// The upstream is a contract fixture, NOT a real Captain model.
func TestTutorRealHTTPTransport(t *testing.T) {
    upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/v1/chat/completions" || r.Header.Get("Authorization") != "Bearer captain-test-key" {
            t.Errorf("invalid Captain upstream request: %s", r.URL.Path)
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        var got struct {
            Model string `json:"model"`
            MaxTokens int `json:"max_tokens"`
            Messages []struct { Content string `json:"content"` } `json:"messages"`
        }
        if err := json.NewDecoder(r.Body).Decode(&got); err != nil { t.Error(err) }
        if got.Model != "aegis-flagship" || got.MaxTokens != 64 || len(got.Messages) != 2 || got.Messages[1].Content != "Explain an API" {
            t.Errorf("unexpected Captain request: %+v", got)
        }
        w.Header().Set("Content-Type", "application/json")
        _, _ = io.WriteString(w, `{"choices":[{"message":{"content":"An API defines a software interface."}}]}`)
    }))
    defer upstream.Close()

    gateway := newTestServer(t)
    gateway.Tutor = TutorGateway{UpstreamURL: upstream.URL, UpstreamKey: "captain-test-key", AcademyKey: "learner-test-key"}
    academy := httptest.NewServer(gateway.handler())
    defer academy.Close()

    call := func(token string, question string) *http.Response {
        t.Helper()
        req, err := http.NewRequest(http.MethodPost, academy.URL+"/api/v1/tutor/chat", strings.NewReader(`{"question":`+question+`}`))
        if err != nil { t.Fatal(err) }
        req.Header.Set("Content-Type", "application/json")
        if token != "" { req.Header.Set("Authorization", "Bearer "+token) }
        resp, err := academy.Client().Do(req)
        if err != nil { t.Fatal(err) }
        return resp
    }
    unauthorized := call("", `"Explain an API"`)
    unauthorized.Body.Close()
    if unauthorized.StatusCode != 401 { t.Fatalf("unauthorized: %d", unauthorized.StatusCode) }

    allowed := call("learner-test-key", `"Explain an API"`)
    defer allowed.Body.Close()
    var answer struct { Answer string `json:"answer"`; Source string `json:"source"` }
    if err := json.NewDecoder(allowed.Body).Decode(&answer); err != nil { t.Fatal(err) }
    if allowed.StatusCode != 200 || answer.Answer != "An API defines a software interface." || answer.Source != "captain-live" {
        t.Fatalf("response %d %+v", allowed.StatusCode, answer)
    }
    if got := allowed.Header.Get("Cache-Control"); got != "no-store" { t.Fatalf("cache policy %q", got) }

    tooLong := call("learner-test-key", `"`+strings.Repeat("x", 4001)+`"`)
    tooLong.Body.Close()
    if tooLong.StatusCode != 400 { t.Fatalf("oversized question: %d", tooLong.StatusCode) }

    upstream.Close()
    unavailable := call("learner-test-key", `"Explain an API"`)
    unavailable.Body.Close()
    if unavailable.StatusCode != 502 { t.Fatalf("unavailable Captain must fail closed: %d", unavailable.StatusCode) }
}
