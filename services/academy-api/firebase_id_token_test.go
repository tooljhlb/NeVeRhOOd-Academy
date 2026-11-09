package main
import (
 "context"
 "crypto"
 "crypto/rand"
 "crypto/rsa"
 "crypto/sha256"
 "encoding/base64"
 "encoding/json"
 "errors"
 "strings"
 "testing"
 "time"
)
func signFirebaseToken(t *testing.T,key *rsa.PrivateKey,claims map[string]any,alg string)string{
 t.Helper();h,_:=json.Marshal(map[string]any{"alg":alg,"kid":"test-key","typ":"JWT"});p,_:=json.Marshal(claims)
 a:=base64.RawURLEncoding.EncodeToString(h)+"."+base64.RawURLEncoding.EncodeToString(p)
 d:=sha256.Sum256([]byte(a));signature,err:=rsa.SignPKCS1v15(rand.Reader,key,crypto.SHA256,d[:]);if err!=nil{t.Fatal(err)}
 return a+"."+base64.RawURLEncoding.EncodeToString(signature)
}
func TestFirebaseVerifierValidAndInvalid(t *testing.T){
 key,err:=rsa.GenerateKey(rand.Reader,2048);if err!=nil{t.Fatal(err)}
 now:=time.Unix(1794180000,0)
 v:=&FirebaseJWTVerifier{ProjectID:"test-project",Now:func()time.Time{return now},Keys:func(context.Context)(map[string]*rsa.PublicKey,error){return map[string]*rsa.PublicKey{"test-key":&key.PublicKey},nil}}
 claims:=func()map[string]any{return map[string]any{"iss":"https://securetoken.google.com/test-project","aud":"test-project","sub":"u-1","exp":now.Unix()+3600,"iat":now.Unix()-60,"auth_time":now.Unix()-360,"email":"test@example.com","email_verified":true}}
 ok:=signFirebaseToken(t,key,claims(),"RS256");id,err:=v.Verify(context.Background(),"Bearer "+ok);if err!=nil||id.UID!="u-1"||!id.EmailVerified{t.Fatalf("valid token: %+v %v",id,err)}
 for _,tc:=range []struct{name string;alter func(map[string]any)}{
  {"wrong_audience",func(c map[string]any){c["aud"]="wrong"}},
  {"wrong_issuer",func(c map[string]any){c["iss"]="https://example.com"}},
  {"expired",func(c map[string]any){c["exp"]=now.Unix()}},
  {"future_issue",func(c map[string]any){c["iat"]=now.Unix()+2}},
  {"future_auth",func(c map[string]any){c["auth_time"]=now.Unix()+2}},
  {"missing_uid",func(c map[string]any){c["sub"]=""}},
 }{t.Run(tc.name,func(t *testing.T){c:=claims();tc.alter(c);token:=signFirebaseToken(t,key,c,"RS256");_,e:=v.Verify(context.Background(),"Bearer "+token);if !errors.Is(e,ErrUnauthorized){t.Fatalf("accepted invalid token: %v",e)}})}
 _,err=v.Verify(context.Background(),"Bearer "+signFirebaseToken(t,key,claims(),"HS256"));if !errors.Is(err,ErrUnauthorized){t.Fatal("alg confusion accepted")}
 _,err=v.Verify(context.Background(),"Bearer "+ok[:len(ok)-2]+"aa");if !errors.Is(err,ErrUnauthorized){t.Fatal("tampered signature accepted")}
 _,err=v.Verify(context.Background(),"Bearer "+ok+" extra");if !errors.Is(err,ErrUnauthorized){t.Fatal("trailing bearer text accepted")}
 _,err=v.Verify(context.Background(),"");if !errors.Is(err,ErrUnauthorized){t.Fatal("missing bearer accepted")}
 _,err=(&FirebaseJWTVerifier{ProjectID:"test-project",Now:func()time.Time{return now},Keys:func(context.Context)(map[string]*rsa.PublicKey,error){return nil,errors.New("down")}}).Verify(context.Background(),"Bearer "+ok);if !errors.Is(err,ErrAuthUnavailable){t.Fatal("key outage not fail closed")}
 if !strings.Contains(ok,"."){t.Fatal("test token malformed")}
}
