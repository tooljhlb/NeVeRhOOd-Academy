package main

import (
 "context"
 "crypto"
 "crypto/rsa"
 "crypto/sha256"
 "crypto/x509"
 "encoding/base64"
 "encoding/json"
 "encoding/pem"
 "errors"
 "io"
 "net/http"
 "strings"
 "sync"
 "time"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrAuthUnavailable = errors.New("identity verification unavailable")

// FirebaseIdentity is obtained exclusively from a cryptographically verified Firebase ID token.
type FirebaseIdentity struct {
 UID string `json:"uid"`
 Email string `json:"email,omitempty"`
 EmailVerified bool `json:"email_verified"`
}

type FirebaseJWTVerifier struct {
 ProjectID string
 Keys func(context.Context) (map[string]*rsa.PublicKey,error)
 Now func() time.Time
}

type firebaseClaims struct {
 Iss string `json:"iss"`
 Aud string `json:"aud"`
 Sub string `json:"sub"`
 Exp int64 `json:"exp"`
 Iat int64 `json:"iat"`
 AuthTime int64 `json:"auth_time"`
 Email string `json:"email"`
 EmailVerified bool `json:"email_verified"`
}

func (v *FirebaseJWTVerifier) Verify(ctx context.Context, authorization string) (FirebaseIdentity,error) {
 deny:=func()(FirebaseIdentity,error){return FirebaseIdentity{},ErrUnauthorized}
 if v==nil||v.ProjectID==""||v.Keys==nil{return FirebaseIdentity{},ErrAuthUnavailable}
 if !strings.HasPrefix(authorization,"Bearer ")||strings.Count(authorization," ")!=1{return deny()}
 token:=strings.TrimPrefix(authorization,"Bearer ")
 if len(token)<30||len(token)>16384{return deny()}
 parts:=strings.Split(token,".");if len(parts)!=3{return deny()}
 rawHeader,err:=base64.RawURLEncoding.DecodeString(parts[0]);if err!=nil||len(rawHeader)>2048{return deny()}
 var header struct{Alg string `json:"alg"`;Kid string `json:"kid"`;Typ string `json:"typ"`}
 if json.Unmarshal(rawHeader,&header)!=nil||header.Alg!="RS256"||header.Kid==""||len(header.Kid)>256{return deny()}
 if header.Typ!=""&&header.Typ!="JWT"{return deny()}
 rawClaims,err:=base64.RawURLEncoding.DecodeString(parts[1]);if err!=nil||len(rawClaims)>8192{return deny()}
 var claims firebaseClaims
 if json.Unmarshal(rawClaims,&claims)!=nil{return deny()}
 now:=time.Now();if v.Now!=nil{now=v.Now()}
 timestamp:=now.Unix()
 if claims.Iss!="https://securetoken.google.com/"+v.ProjectID||claims.Aud!=v.ProjectID{return deny()}
 if len(claims.Sub)<1||len(claims.Sub)>128||claims.Exp<=timestamp||claims.Iat<=0||claims.Iat>timestamp||claims.AuthTime<=0||claims.AuthTime>timestamp{return deny()}
 sig,err:=base64.RawURLEncoding.DecodeString(parts[2]);if err!=nil{return deny()}
 keys,err:=v.Keys(ctx);if err!=nil{return FirebaseIdentity{},ErrAuthUnavailable}
 key:=keys[header.Kid];if key==nil{return deny()}
 digest:=sha256.Sum256([]byte(parts[0]+"."+parts[1]))
 if rsa.VerifyPKCS1v15(key,crypto.SHA256,digest[:],sig)!=nil{return deny()}
 return FirebaseIdentity{UID:claims.Sub,Email:claims.Email,EmailVerified:claims.EmailVerified},nil
}

// GoogleFirebaseCertCache is concurrency-safe and fails closed if Google certificate retrieval fails.
type GoogleFirebaseCertCache struct {
 mu sync.Mutex
 until time.Time
 keys map[string]*rsa.PublicKey
 client *http.Client
}

func NewGoogleFirebaseCertCache() *GoogleFirebaseCertCache {return &GoogleFirebaseCertCache{client:&http.Client{Timeout:5*time.Second}}}
func (c *GoogleFirebaseCertCache) Keys(ctx context.Context)(map[string]*rsa.PublicKey,error){
 c.mu.Lock();defer c.mu.Unlock()
 if time.Now().Before(c.until)&&len(c.keys)>0{return c.keys,nil}
 req,err:=http.NewRequestWithContext(ctx,http.MethodGet,"https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com",nil);if err!=nil{return nil,err}
 resp,err:=c.client.Do(req);if err!=nil{return nil,err};defer resp.Body.Close()
 if resp.StatusCode!=http.StatusOK{return nil,ErrAuthUnavailable}
 data,err:=io.ReadAll(io.LimitReader(resp.Body,1024*1024+1));if err!=nil||len(data)>1024*1024{return nil,ErrAuthUnavailable}
 var certs map[string]string;if json.Unmarshal(data,&certs)!=nil||len(certs)==0{return nil,ErrAuthUnavailable}
 parsed:=make(map[string]*rsa.PublicKey,len(certs))
 for kid,pemString:=range certs{
  block,rest:=pem.Decode([]byte(pemString));if block==nil||len(strings.TrimSpace(string(rest)))>0{return nil,ErrAuthUnavailable}
  cert,err:=x509.ParseCertificate(block.Bytes);if err!=nil{return nil,ErrAuthUnavailable}
  rsaKey,ok:=cert.PublicKey.(*rsa.PublicKey);if !ok{return nil,ErrAuthUnavailable};parsed[kid]=rsaKey
 }
 c.keys=parsed;c.until=time.Now().Add(15*time.Minute)
 return c.keys,nil
}
