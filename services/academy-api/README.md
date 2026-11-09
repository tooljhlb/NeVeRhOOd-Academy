# NeVeRhOOd Academy API (independent Go service)

A stdlib-only HTTP API with a read-only public Academy catalog and a separate **disabled-by-default, authenticated** Captain tutor proxy. The tutor is not publicly hosted. It **does not** contain personal cloud credentials, private mailboxes, model weights or GPU controls.

## Build and run

Requires Go 1.23+ from the repository root:

```sh
cd services/academy-api
go test ./...
go build -o academy-api .
cd ../..
./services/academy-api/academy-api -catalog data/courses.csv -addr 127.0.0.1:8080
```

For remote deployment bind to the host/port provided by an approved platform, e.g. `-addr 0.0.0.0:8080`, and place TLS, access logs and request limits at the reverse proxy. Do **not** expose the owner workstation directly.

## Endpoints

| HTTP | Route | Behavior |
|---|---|---|
| GET | /api/v1/health | Public service health and count |
| GET | /api/v1/session | Authenticated Firebase ID token inspection; 401 invalid, 503 if unconfigured; returns free/not_configured until paid entitlements exist |
| GET | /api/v1/courses | Filter by `query` or `q`, `category`, `verified_only`; `limit` 1..100 and `offset` |
| GET | /api/v1/courses/{id} | Single course record |
| GET | /api/v1/learning-paths | Original three-lesson Foundations track, with certificate status explicitly `not_issued` |
| GET | /api/v1/models | Public availability remains `not_deployed` until a stable, externally verified Captain endpoint exists |
| POST | /api/v1/tutor/chat | **Private, opt-in** authenticated tutor gateway; requires real upstream model, no fake fallback |

Use `ACADEMY_ALLOWED_ORIGIN=https://your-manus-domain` to allow the frontend's browser origin. The default is **no cross-origin access**. Expose no private admin or job endpoints until authentication and authorization exist.

This service reads the currently reviewed public pilot CSV; it does **not** serve the owner's 1,000-entry unverified private research archive. The Academy cannot issue accredited degree credits.

## Firebase ID-token verification (implemented, opt-in)
Set `ACADEMY_FIREBASE_PROJECT_ID=<actual-firebase-project-id>` on the protected server **after** configuring a real Firebase Google sign-in provider and authorized domain. The Go backend verifies signed RS256 Firebase ID tokens and checks their project issuer/audience, issue/auth timestamps, expiry, signature and subject UID against Google public certificates. Never accept an email address or a client-provided subscription role as authentication. The session endpoint returns only a FREE, not-configured membership while Stripe and database entitlements are still unimplemented.

`GET /api/v1/session` requires `Authorization: Bearer <Firebase ID token>` and returns `Cache-Control: no-store`. Without the server configuration it fails closed with 503. This is a **real JWT validation module with local signed-token tests**, not evidence of live Google account integration or persisted learner records.

Build gate: `go test -race -count=1 ./...`, `go vet ./...`, `go build .`. Verified 2026-11-08 against the complete Go package (22 top-level tests).

## Future integration

The Manus client can first use static `site/courses.json`, then move to this API when hosted. Tutor/chat now has a deliberately private gateway in `tutor.go` and `tutor_test.go`. It is **disabled unless all three server environment variables are supplied**:

```sh
CAPTAIN_UPSTREAM_URL=https://your-authenticated-captain-origin
CAPTAIN_SERVER_API_KEY=<backend-bearer-token>
ACADEMY_TUTOR_ACCESS_KEY=<separate-academy-client-bearer-token>
```

For a local isolated demonstration the origin may instead be `http://127.0.0.1:<port>`; other unencrypted HTTP origins are refused. The route accepts `POST /api/v1/tutor/chat` with `Authorization: Bearer <academy-client-token>` and JSON `{ "question": "Explain APIs" }`. It forwards a **bounded** request to the authenticated Captain `/v1/chat/completions`, applies a 25-second deadline, validates the result, and returns an answer or an explicit 502/503—not a canned reply. It never places the upstream Captain secret in the browser.

**Do not put the Academy client token in static JavaScript.** This endpoint remains an operator/private backend contract until Academy user login, per-session authorization, rate limiting, abuse controls, HTTPS termination and reliable Captain hosting are built. The existing `/api/v1/models` public status remains `not_deployed`. Integration testing against a real T4 model already passed on the native Captain side; that is not evidence of public availability.

Run `go test ./...` and the repository's Python validator before deploying. The same dataset must be shipped together with the service binary.
