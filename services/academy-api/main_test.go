package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func newTestServer(t *testing.T) *Server {
	t.Helper()
	c, err := loadCSV(filepath.Join("..", "..", "data", "courses.csv"))
	if err != nil { t.Fatal(err) }
	return &Server{Courses: c, Origin: "https://academy.example"}
}
func TestHealth(t *testing.T) {
	w:=httptest.NewRecorder()
	newTestServer(t).handler().ServeHTTP(w,httptest.NewRequest("GET","/api/v1/health",nil))
	if w.Code!=200 { t.Fatalf("status %d",w.Code) }
}
func TestSearchAndPagination(t *testing.T) {
	w:=httptest.NewRecorder()
	newTestServer(t).handler().ServeHTTP(w,httptest.NewRequest("GET","/api/v1/courses?q=Harvard&limit=1",nil))
	if w.Code!=200 { t.Fatalf("status %d",w.Code) }
	var result struct { Total int `json:"total"`; Items []Course `json:"items"` }
	if err:=json.Unmarshal(w.Body.Bytes(),&result);err!=nil{t.Fatal(err)}
	if result.Total<1 || len(result.Items)!=1 {t.Fatalf("unexpected result %+v",result)}
}
func TestNoBogusModelEndpoint(t *testing.T) {
	w:=httptest.NewRecorder()
	newTestServer(t).handler().ServeHTTP(w,httptest.NewRequest("GET","/api/v1/models",nil))
	var result map[string]any
	if err:=json.Unmarshal(w.Body.Bytes(),&result);err!=nil{t.Fatal(err)}
	if w.Code!=200 {t.Fatal(w.Code)}
	if string(w.Body.Bytes())=="" {t.Fatal("empty status")}
}
func TestRejectMutations(t *testing.T) {
	w:=httptest.NewRecorder()
	newTestServer(t).handler().ServeHTTP(w,httptest.NewRequest("POST","/api/v1/courses",nil))
	if w.Code!=http.StatusMethodNotAllowed {t.Fatalf("expected 405 got %d",w.Code)}
}
func TestOriginAllowlist(t *testing.T) {
	w:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/api/v1/health",nil)
	req.Header.Set("Origin","https://evil.example")
	newTestServer(t).handler().ServeHTTP(w,req)
	if got:=w.Header().Get("Access-Control-Allow-Origin");got!=""{t.Fatalf("unexpected CORS %q",got)}
}
func TestVerifiedOnlyExcludesUnreviewed(t *testing.T) {
	w:=httptest.NewRecorder()
	newTestServer(t).handler().ServeHTTP(w,httptest.NewRequest("GET","/api/v1/courses?verified_only=true",nil))
	if w.Code!=200{t.Fatalf("unexpected status %d",w.Code)}
	var r struct{ Total int `json:"total"` }
	if err:=json.Unmarshal(w.Body.Bytes(),&r);err!=nil{t.Fatal(err)}
	if r.Total!=0{t.Fatalf("pilot entries are not link-verified; got %d",r.Total)}
}
