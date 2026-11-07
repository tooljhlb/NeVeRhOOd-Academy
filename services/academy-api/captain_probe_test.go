package main

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCaptainNotConfigured(t *testing.T){
    status:=probeCaptain(context.Background(),"")
    if status.State!="not_configured"||status.Ready {t.Fatalf("%+v",status)}
}
func TestCaptainRejectsUnsafeEndpoint(t *testing.T){
    for _,s:=range []string{
        "file:///tmp/ready", "http://example.org/ready",
        "http://127.0.0.1:8000/chat","https://name:password@example.org/ready",
        "https://example.org/ready?secret=test",
    } {
        x:=probeCaptain(context.Background(),s)
        if x.State!="invalid_configuration"||x.Ready {t.Fatalf("accepted %s: %+v",s,x)}
    }
}
func TestCaptainNotReady(t *testing.T){
    s:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        if r.URL.Path!="/ready"{t.Errorf("wrong probe path %s",r.URL.Path)}
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusServiceUnavailable)
        _,_=w.Write([]byte(`{"ready":false}`))
    }))
    defer s.Close()
    x:=probeCaptain(context.Background(),s.URL)
    if x.State!="not_ready"||x.Ready{t.Fatalf("%+v",x)}
}
func TestCaptainHealthyRequiresPositiveJSON(t *testing.T){
    s:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        w.Header().Set("Content-Type","application/json")
        _,_=w.Write([]byte(`{"ready":true}`))
    }))
    defer s.Close()
    x:=probeCaptain(context.Background(),s.URL)
    if x.State!="ready"||!x.Ready{t.Fatalf("%+v",x)}
}
func TestCaptainMalformattedPositiveHTTPIsNotReady(t *testing.T){
    s:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        _,_=w.Write([]byte("not json"))
    }))
    defer s.Close()
    x:=probeCaptain(context.Background(),s.URL)
    if x.State!="invalid_response"||x.Ready{t.Fatalf("%+v",x)}
}
func TestCaptainReadinessAPI(t *testing.T){
    s:=newTestServer(t)
    w:=httptest.NewRecorder()
    s.handler().ServeHTTP(w,httptest.NewRequest("GET","/api/v1/captain/status",nil))
    if w.Code!=200{t.Fatalf("status %d",w.Code)}
    var info CaptainReadiness
    if err:=json.Unmarshal(w.Body.Bytes(),&info);err!=nil{t.Fatal(err)}
    if info.Ready||info.State!="not_configured"{t.Fatalf("%+v",info)}
}
