package main

import (
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "strings"
 "testing"
)

func tutorRequest(t *testing.T, gw TutorGateway, question string, key string) *httptest.ResponseRecorder {
 t.Helper()
 body,_:=json.Marshal(map[string]string{"question":question})
 req:=httptest.NewRequest(http.MethodPost,"/api/v1/tutor/chat",strings.NewReader(string(body)))
 if key!="" {req.Header.Set("Authorization","Bearer "+key)}
 s:=newTestServer(t)
 s.Tutor=gw
 out:=httptest.NewRecorder()
 s.handler().ServeHTTP(out,req)
 return out
}
func TestTutorDefaultDisabled(t *testing.T) {
 if w:=tutorRequest(t,TutorGateway{},"Explain APIs","");w.Code!=503 {t.Fatalf("status %d",w.Code)}
}
func TestTutorRejectsWithoutAcademyBearer(t *testing.T) {
 gw:=TutorGateway{UpstreamURL:"http://127.0.0.1:9876",UpstreamKey:"origin",AcademyKey:"learner"}
 if w:=tutorRequest(t,gw,"hi","");w.Code!=401 {t.Fatalf("status %d",w.Code)}
}
func TestTutorRejectsInvalidQuestion(t *testing.T) {
 gw:=TutorGateway{UpstreamURL:"http://127.0.0.1:9876",UpstreamKey:"origin",AcademyKey:"learner"}
 if w:=tutorRequest(t,gw,"","learner");w.Code!=400 {t.Fatalf("status %d",w.Code)}
}
func TestTutorRejectsUntrustedOrigin(t *testing.T) {
 for _,v:=range []string{"http://example.com","https://user:pw@example.com","https://example.com/path","https://example.com/?secret=1"} {
  if _,err:=validateCaptainOrigin(v);err==nil {t.Fatalf("accepted %q",v)}
 }
}
func TestTutorToAuthenticatedCaptain(t *testing.T) {
 mock:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
  if r.Header.Get("Authorization")!="Bearer origin" {t.Errorf("wrong origin credentials")}
  var payload struct{Messages []struct{Content string};MaxTokens int}
  if err:=json.NewDecoder(r.Body).Decode(&payload);err!=nil {t.Error(err)}
  if len(payload.Messages)!=2 || payload.Messages[1].Content!="Explain APIs" || payload.MaxTokens!=64 {t.Errorf("payload %+v",payload)}
  w.Header().Set("Content-Type","application/json")
  _,_=w.Write([]byte("{\"choices\":[{\"message\":{\"content\":\"An API defines an interface.\"}}]}"))
 }))
 defer mock.Close()
 gw:=TutorGateway{UpstreamURL:mock.URL,UpstreamKey:"origin",AcademyKey:"learner"}
 w:=tutorRequest(t,gw,"Explain APIs","learner")
 if w.Code!=200 {t.Fatalf("status %d: %s",w.Code,w.Body.String())}
 if !strings.Contains(w.Body.String(),"An API defines an interface.") {t.Fatalf("no answer %s",w.Body.String())}
 if w.Header().Get("Cache-Control")!="no-store" {t.Fatal("tutor content cached")}
}
func TestTutorRefusesFakeFallbackOnBadUpstream(t *testing.T) {
 mock:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.WriteHeader(503)}))
 defer mock.Close()
 gw:=TutorGateway{UpstreamURL:mock.URL,UpstreamKey:"origin",AcademyKey:"learner"}
 if w:=tutorRequest(t,gw,"Hello","learner");w.Code!=502 {t.Fatalf("status %d",w.Code)}
}
