// NeVeRhOOd Academy read-only API. No owner credentials or private compute dependencies.
package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Course struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Provider         string `json:"provider"`
	Category         string `json:"category"`
	URL              string `json:"url"`
	CredentialStatus string `json:"credential_status"`
	LinkStatus       string `json:"link_status"`
}

type Server struct {
	Courses []Course
	Origin  string
	CaptainProbeURL string
	Tutor TutorGateway
}

func loadCSV(path string) ([]Course, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil { return nil, err }
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil { return nil, fmt.Errorf("read CSV: %w", err) }
	if len(rows) < 2 { return nil, errors.New("catalog empty") }
	keys := map[string]int{}
	for i, field := range rows[0] { keys[strings.TrimSpace(field)] = i }
	required := []string{"id", "title", "provider", "category", "url", "credential_status", "link_status"}
	for _, key := range required {
		if _, ok := keys[key]; !ok { return nil, fmt.Errorf("missing CSV field %q", key) }
	}
	get := func(row []string, key string) string {
		i := keys[key]
		if i >= len(row) { return "" }
		return strings.TrimSpace(row[i])
	}
	seen := map[string]bool{}
	courses := make([]Course, 0, len(rows)-1)
	for i, row := range rows[1:] {
		c := Course{
			ID: get(row, "id"), Title: get(row, "title"),
			Provider: get(row, "provider"), Category: get(row, "category"),
			URL: get(row, "url"), CredentialStatus: get(row, "credential_status"),
			LinkStatus: get(row, "link_status"),
		}
		if c.ID == "" || c.Title == "" || !strings.HasPrefix(c.URL, "https://") {
			return nil, fmt.Errorf("invalid catalog row %d", i+2)
		}
		if seen[c.ID] { return nil, fmt.Errorf("duplicate course id %q", c.ID) }
		seen[c.ID] = true
		courses = append(courses, c)
	}
	return courses, nil
}

func jsonReply(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (s *Server) handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/tutor/chat", s.Tutor)
	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		jsonReply(w, http.StatusOK, map[string]any{
			"status": "ok", "service": "neverhood-academy-api",
			"catalog_records": len(s.Courses), "api_version": "v1",
		})
	})
	mux.HandleFunc("/api/v1/captain/status", func(w http.ResponseWriter, r *http.Request) {
        jsonReply(w, http.StatusOK, probeCaptain(r.Context(), s.CaptainProbeURL))
    })
    mux.HandleFunc("/api/v1/models", func(w http.ResponseWriter, r *http.Request) {
		jsonReply(w, http.StatusOK, map[string]any{
			"models": []any{map[string]any{"id": "captain-aegis", "status": "not_deployed",
				"chat_endpoint": nil, "note": "The real Captain backend requires GPU and full checkpoint validation."}},
		})
	})
	mux.HandleFunc("/api/v1/learning-paths", func(w http.ResponseWriter, r *http.Request) {
        jsonReply(w, http.StatusOK, map[string]any{"items": []any{
            map[string]any{"id":"ai-engineering-foundations","title":"AI Engineering Foundations","accredited":false,"lessons":[]any{
                map[string]string{"id":"ai-agents-101","title":"AI Agents 101","url_path":"/learn/ai-agents-101/index.html"},
                map[string]string{"id":"protocols-101","title":"Protocols 101","url_path":"/learn/protocols-101/index.html"},
                map[string]string{"id":"evaluation-101","title":"Model Evaluation 101","url_path":"/learn/evaluation-101/index.html"},
            },"certificate_status":"not_issued","url_path":"/track/foundations/index.html"},
        }})
    })
    mux.HandleFunc("/api/v1/courses", func(w http.ResponseWriter, r *http.Request) {
		q := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("query")))
		if q == "" { q = strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q"))) }
		category := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("category")))
		verifiedOnly := r.URL.Query().Get("verified_only") == "true"
		limit := 50
		if raw := r.URL.Query().Get("limit"); raw != "" {
			v, e := strconv.Atoi(raw)
			if e != nil || v < 1 || v > 100 { jsonReply(w, 400, map[string]string{"error": "limit must be 1..100"}); return }
			limit = v
		}
		offset := 0
		if raw := r.URL.Query().Get("offset"); raw != "" {
			v, e := strconv.Atoi(raw)
			if e != nil || v < 0 { jsonReply(w, 400, map[string]string{"error": "offset must be >=0"}); return }
			offset = v
		}
		matches := make([]Course, 0)
		for _, c := range s.Courses {
			if verifiedOnly && c.LinkStatus != "verified" { continue }
			if q != "" && !strings.Contains(strings.ToLower(c.Title+" "+c.Provider+" "+c.Category), q) { continue }
			if category != "" && !strings.EqualFold(c.Category, category) { continue }
			matches = append(matches, c)
		}
		total := len(matches)
		if offset > total { offset = total }
		end := offset+limit
		if end > total { end = total }
		jsonReply(w, 200, map[string]any{"items": matches[offset:end], "total": total, "offset": offset, "limit": limit})
	})
	mux.HandleFunc("/api/v1/courses/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/v1/courses/")
		if id == "" || strings.Contains(id, "/") { jsonReply(w, 404, map[string]string{"error": "not found"}); return }
		for _, c := range s.Courses { if c.ID == id { jsonReply(w, 200, c); return } }
		jsonReply(w, 404, map[string]string{"error": "course not found"})
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		if s.Origin != "" && r.Header.Get("Origin") == s.Origin {
			w.Header().Set("Access-Control-Allow-Origin", s.Origin)
			w.Header().Set("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			if s.Origin == "" || r.Header.Get("Origin") != s.Origin { w.WriteHeader(403); return }
			w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS, POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodGet && r.Method != http.MethodHead && !(r.Method == http.MethodPost && r.URL.Path == "/api/v1/tutor/chat") {
			w.Header().Set("Allow", "GET, HEAD, OPTIONS")
			jsonReply(w, http.StatusMethodNotAllowed, map[string]string{"error": "read-only API"})
			return
		}
		mux.ServeHTTP(w,r)
	})
}

func main() {
	catalog := flag.String("catalog", "data/courses.csv", "Academy CSV dataset (verified status kept per entry)")
	addr := flag.String("addr", "127.0.0.1:8080", "HTTP bind address")
	flag.Parse()
	courses, err := loadCSV(*catalog)
	if err != nil { log.Fatal(err) }
	origin := strings.TrimSpace(os.Getenv("ACADEMY_ALLOWED_ORIGIN"))
	s := &Server{Courses: courses, Origin: origin, CaptainProbeURL: strings.TrimSpace(os.Getenv("CAPTAIN_READY_URL")), Tutor: TutorGateway{UpstreamURL: strings.TrimSpace(os.Getenv("CAPTAIN_UPSTREAM_URL")), UpstreamKey: os.Getenv("CAPTAIN_SERVER_API_KEY"), AcademyKey: os.Getenv("ACADEMY_TUTOR_ACCESS_KEY")}}
	server := &http.Server{
		Addr: *addr, Handler: s.handler(), ReadHeaderTimeout: 5*time.Second,
		ReadTimeout: 15*time.Second, WriteTimeout: 30*time.Second, IdleTimeout: 60*time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = server.Shutdown(c)
	}()
	log.Printf("Academy API listening on %s with %d catalog records", *addr, len(courses))
	if err := server.ListenAndServe(); err != nil && !errors.Is(err,http.ErrServerClosed) { log.Fatal(err) }
}
