package main

import (
  "context"
  "crypto/subtle"
  "encoding/json"
  "errors"
  "io"
  "net"
  "net/http"
  "net/url"
  "strings"
  "time"
)

// TutorGateway is an opt-in backend adapter; no fake answers or public tokens.
type TutorGateway struct {
  UpstreamURL string
  UpstreamKey string
  AcademyKey string
  Model string
  Client *http.Client
}
func (g TutorGateway) configured() bool {
  return g.UpstreamURL != "" && g.UpstreamKey != "" && g.AcademyKey != ""
}
func tutorResponse(w http.ResponseWriter, status int, body any) {
  w.Header().Set("Content-Type","application/json; charset=utf-8")
  w.Header().Set("Cache-Control","no-store")
  w.Header().Set("X-Content-Type-Options","nosniff")
  w.WriteHeader(status)
  _ = json.NewEncoder(w).Encode(body)
}
func validateCaptainOrigin(raw string) (string,error) {
  u,err := url.Parse(raw)
  if err!=nil || u==nil || u.User!=nil || u.Host=="" || u.RawQuery!="" || u.Fragment!="" || u.Path!="" && u.Path!="/" {
    return "",errors.New("invalid upstream origin")
  }
  host:=u.Hostname()
  if u.Scheme!="https" {
    ip:=net.ParseIP(host)
    if u.Scheme!="http" || !(host=="localhost" || ip!=nil && ip.IsLoopback()) {
      return "",errors.New("upstream must use HTTPS or loopback HTTP")
    }
  }
  return strings.TrimRight(u.String(),"/")+"/v1/chat/completions",nil
}
func (g TutorGateway) ServeHTTP(w http.ResponseWriter,r *http.Request) {
  if r.Method!=http.MethodPost {
    w.Header().Set("Allow","POST")
    tutorResponse(w,405,map[string]string{"error":"POST only"})
    return
  }
  if !g.configured() {
    tutorResponse(w,503,map[string]string{"error":"Captain tutor not configured"})
    return
  }
  auth:=r.Header.Get("Authorization")
  token:=strings.TrimPrefix(auth,"Bearer ")
  if !strings.HasPrefix(auth,"Bearer ") || subtle.ConstantTimeCompare([]byte(token),[]byte(g.AcademyKey))!=1 {
    tutorResponse(w,401,map[string]string{"error":"Unauthorized"})
    return
  }
  endpoint,err:=validateCaptainOrigin(g.UpstreamURL)
  if err!=nil {
    tutorResponse(w,503,map[string]string{"error":"Captain origin invalid"})
    return
  }
  r.Body=http.MaxBytesReader(w,r.Body,8192)
  var input struct { Question string }
  decoder:=json.NewDecoder(r.Body)
  decoder.DisallowUnknownFields()
  if err:=decoder.Decode(&input);err!=nil {
    tutorResponse(w,400,map[string]string{"error":"Invalid JSON question"})
    return
  }
  var extra any
  if decoder.Decode(&extra)!=io.EOF {
    tutorResponse(w,400,map[string]string{"error":"Trailing JSON"})
    return
  }
  question:=strings.TrimSpace(input.Question)
  if question=="" || len(question)>4000 {
    tutorResponse(w,400,map[string]string{"error":"Question must be 1..4000 bytes"})
    return
  }
  model:=g.Model
  if model=="" {model="aegis-flagship"}
  payload:=map[string]any{
    "model":model,"max_tokens":128,"temperature":0.2,"stream":false,
    "messages":[]map[string]string{
      {"role":"system","content":"You are a study tutor. Give grounded explanations, distinguish uncertainty, and never claim accredited credentials or course completion."},
      {"role":"user","content":question},
    },
  }
  serialized,err:=json.Marshal(payload)
  if err!=nil {tutorResponse(w,500,map[string]string{"error":"Serialization failed"});return}
  ctx,cancel:=context.WithTimeout(r.Context(),25*time.Second)
  defer cancel()
  req,err:=http.NewRequestWithContext(ctx,http.MethodPost,endpoint,strings.NewReader(string(serialized)))
  if err!=nil {tutorResponse(w,503,map[string]string{"error":"Captain request invalid"});return}
  req.Header.Set("Authorization","Bearer "+g.UpstreamKey)
  req.Header.Set("Content-Type","application/json")
  client:=g.Client
  if client==nil {
    client=&http.Client{Timeout:25*time.Second,CheckRedirect:func(*http.Request,[]*http.Request)error{return http.ErrUseLastResponse}}
  }
  resp,err:=client.Do(req)
  if err!=nil {tutorResponse(w,502,map[string]string{"error":"Captain unavailable"});return}
  defer resp.Body.Close()
  if resp.StatusCode!=200 {tutorResponse(w,502,map[string]string{"error":"Captain rejected request"});return}
  var output struct{Choices []struct{Message struct{Content string}}}
  if err:=json.NewDecoder(io.LimitReader(resp.Body,65537)).Decode(&output);err!=nil || len(output.Choices)==0 || strings.TrimSpace(output.Choices[0].Message.Content)=="" {
    tutorResponse(w,502,map[string]string{"error":"Invalid Captain response"})
    return
  }
  tutorResponse(w,200,map[string]any{"answer":output.Choices[0].Message.Content,"model":model,"source":"captain-live","credential_status":"not_accredited"})
}
