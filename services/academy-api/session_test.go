package main

import (
 "context"
 "crypto/rand"
 "crypto/rsa"
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "strings"
 "testing"
 "time"
)

// Auth contracts are exercised without any live Google credentials or merchant account.
func TestAcademySessionAuthentication(t *testing.T) {
 s:=&Server{}
 call:=func(authorization string)*httptest.ResponseRecorder{
  req:=httptest.NewRequest(http.MethodGet,"/api/v1/session",nil)
  if authorization!="" {req.Header.Set("Authorization",authorization)}
  w:=httptest.NewRecorder()
  s.handler().ServeHTTP(w,req)
  return w
 }
 if w:=call("");w.Code!=503{t.Fatalf("missing auth integration should fail closed with 503: %d",w.Code)}
 key,err:=rsa.GenerateKey(rand.Reader,2048);if err!=nil{t.Fatal(err)}
 now:=time.Unix(1794180000,0)
 s.Firebase=&FirebaseJWTVerifier{ProjectID:"academy-example",Now:func()time.Time{return now},Keys:func(context.Context)(map[string]*rsa.PublicKey,error){return map[string]*rsa.PublicKey{"test-key":&key.PublicKey},nil}}
 if w:=call("");w.Code!=401{t.Fatalf("missing bearer should be 401: %d",w.Code)}
 claims:=map[string]any{"iss":"https://securetoken.google.com/academy-example","aud":"academy-example","sub":"learner-a","exp":now.Unix()+120,"iat":now.Unix()-120,"auth_time":now.Unix()-240,"email":"learner@example.com","email_verified":true}
 token:=signFirebaseToken(t,key,claims,"RS256")
 w:=call("Bearer "+token)
 if w.Code!=http.StatusOK{t.Fatalf("valid Firebase bearer was rejected: %d %s",w.Code,w.Body.String())}
 if w.Header().Get("Cache-Control")!="no-store"{t.Fatal("private session response must not be cached")}
 var output struct{UID string `json:"uid"`;Plan string `json:"plan"`;SubscriptionStatus string `json:"subscription_status"`}
 if err:=json.NewDecoder(strings.NewReader(w.Body.String())).Decode(&output);err!=nil{t.Fatal(err)}
 if output.UID!="learner-a"||output.Plan!="free"||output.SubscriptionStatus!="not_configured"{t.Fatalf("unexpected session %v",output)}
 claims["aud"]="somebody-else";bad:=signFirebaseToken(t,key,claims,"RS256")
 if w:=call("Bearer "+bad);w.Code!=401{t.Fatalf("invalid audience accepted: %d",w.Code)}
}
