package main

import (
    "context"
    "encoding/json"
    "errors"
    "net/http"
    "net/url"
    "strings"
    "time"
)

// CaptainReadiness is deliberately metadata-only. It never loads weights, stores
// credentials or proxies chat. The model service runs independently on a GPU host.
type CaptainReadiness struct {
    Model string `json:"model"`
    State string `json:"state"`
    Ready bool `json:"ready"`
    Note  string `json:"note"`
}

func probeCaptain(ctx context.Context, target string) CaptainReadiness {
    status := CaptainReadiness{Model:"captain-aegis",State:"not_configured",Note:"No independently served Captain inference endpoint has been configured."}
    if strings.TrimSpace(target)=="" {return status}
    u,err:=url.Parse(target)
    if err!=nil || u.User!=nil || u.RawQuery!="" || u.Fragment!="" || u.Host=="" {
        status.State="invalid_configuration";status.Note="Captain readiness address is invalid.";return status
    }
    host:=strings.ToLower(u.Hostname())
    // Plain HTTP is restricted to local loopback. Public GPU services require HTTPS.
    if u.Scheme!="https" && !(u.Scheme=="http"&&(host=="127.0.0.1"||host=="localhost"||host=="::1")) {
        status.State="invalid_configuration";status.Note="Only HTTPS or loopback HTTP is permitted.";return status
    }
    if u.Path!=""&&u.Path!="/"&&u.Path!="/ready" {
        status.State="invalid_configuration";status.Note="Configure only the Captain /ready endpoint.";return status
    }
    u.Path="/ready"
    req,err:=http.NewRequestWithContext(ctx,http.MethodGet,u.String(),nil)
    if err!=nil {status.State="invalid_configuration";return status}
    client:=&http.Client{Timeout:2*time.Second,CheckRedirect:func(*http.Request,[]*http.Request)error{return http.ErrUseLastResponse}}
    resp,err:=client.Do(req)
    if err!=nil {
        status.State="offline";status.Note="Captain readiness probe failed; inference is not claimed live.";return status
    }
    defer resp.Body.Close()
    if resp.StatusCode!=200 {
        status.State="not_ready";status.Note="Captain backend did not pass its readiness gate.";return status
    }
    var body struct { Ready bool `json:"ready"` }
    if err=json.NewDecoder(http.MaxBytesReader(nil,resp.Body,4096)).Decode(&body);err!=nil&& !errors.Is(err,context.Canceled){
        status.State="invalid_response";status.Note="Captain readiness response could not be validated.";return status
    }
    if !body.Ready {
        status.State="not_ready";status.Note="Captain backend has not loaded its model.";return status
    }
    status.State="ready";status.Ready=true;status.Note="Backend reported loaded; chat quality and authorization require separate tests."
    return status
}
