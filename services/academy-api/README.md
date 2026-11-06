# NeVeRhOOd Academy API (independent Go service)

A stdlib-only, read-only HTTP API for the public Academy catalog and Manus frontend. It **does not** contain personal cloud credentials, private mailboxes, model weights or GPU controls.

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
| GET | /api/v1/courses | Filter by `query` or `q`, `category`, `verified_only`; `limit` 1..100 and `offset` |
| GET | /api/v1/courses/{id} | Single course record |
| GET | /api/v1/models | Captain remains `not_deployed` until real inference is verified |

Use `ACADEMY_ALLOWED_ORIGIN=https://your-manus-domain` to allow the frontend's browser origin. The default is **no cross-origin access**. Expose no private admin or job endpoints until authentication and authorization exist.

This service reads the currently reviewed public pilot CSV; it does **not** serve the owner's 1,000-entry unverified private research archive. The Academy cannot issue accredited degree credits.

## Future integration

The Manus client can first use static `site/courses.json`, then move to this API when hosted. Tutor/chat will be a separate authenticated, rate-limited Captain service, not a fake chatbot in this API.

Run `go test ./...` and the repository's Python validator before deploying. The same dataset must be shipped together with the service binary.
