# Academy release engineering receipt — 2026-11-08

**Evidence status:** PASS for static source, native API testing and portable package integrity. **NOT** production launch or Google/Manus account activation.

## Source
- Academy main revision verified: `b05f708e2dd9ef8f1ec4c71344d6e69f62e31f7d`.
- Publishing branch `gh-pages` revision: `74aaaf66680f54ced0189c9f3a85b6bb9b1c25fd`.
- Hosted Pages URL not yet verified. The Pages source may require repository Settings → Pages → Deploy from branch → gh-pages → /(root), and the account's earlier GitHub billing/Actions restriction still requires resolution. No QR link should use a broken/unverified URL.
- Static Manus handoff Drive file: `NEVERHOOD_ACADEMY_MANUS_HANDOFF_20261108.zip`; 52 files; 106,226 bytes; SHA256 `fe7384f5c0527f97e8caf1d19209e1522dd1f059dcddd5dcf056bf5b27ae8edd`; archive `testzip` PASS. Download URL: https://drive.google.com/file/d/1RjXquAgt8uQ6-DimupZ7XNQZYXQekcKb/view?usp=drivesdk

## Tested commands and observed results
| Step | Command | Actual result |
| --- | --- | --- |
| Academy Python | `python -m unittest discover -s tests` | 13 tests, PASS |
| Workspace and metric lab Node | `node --test tests/workspace.test.cjs tests/evaluation_lab.test.cjs` | 12 JavaScript tests, PASS |
| Academy static builder | `python scripts/build.py` | 4 unique course records built, PASS |
| Native Go API | `cd services/academy-api && go test -race -count=1 ./...` | PASS; 22 top-level tests including Firebase verifier/session HTTP tests |
| Go vet | `go vet ./...` | PASS |
| Go binary compilation | `go build -o /tmp/academy-api .` | PASS |
| Model GPU proof | Earlier isolated Academy Go ↔ Captain FastAPI ↔ Tesla T4 E2E | PASS in previous receipt; not repeated and NOT 24/7 hosting |

## Added since previous handoff
- Stable responsive UI navigation with keyboard and focus management; Studies/Projects/Models tabs; safe lesson-viewer allowlist.
- Firebase JWT verification from Google-authenticated RSA public certs, strict issuer/audience/sub/iat/auth_time/expiry and signature checks; fails closed if unconfigured, invalid or cert service unavailable.
- `GET /api/v1/session`: server-side user identity from JWT. Returns `subscription_status=not_configured` and `plan=free` until proper database/billing integration.
- API-level tests for nil auth setup (503), missing token (401), valid token (200), invalid audience (401) and no-store response cache policy.

## Security boundaries
- No customer Google login is live; `ACADEMY_FIREBASE_PROJECT_ID` and authorized OAuth domain are not configured in a deployed application.
- No student progress persistence, paid Stripe checkout/webhooks or production model provider billing yet.
- Existing separate `/api/v1/tutor/chat` shared-secret pathway remains operator-only; do NOT expose publicly until per-user session auth/entitlements and quotas replace it.
- The user-facing chat send button is disabled until a real provider is integrated and metered.
- Hugging Face public model card available; no 24/7 Captain hosted inference claim. Aegis Data Frameworks private worker remains isolated.

## Launch readiness
The owner public-launch weighted scoreboard stays **55/100**. Verified new code reduces engineering backlog, but remaining points require all gate-specific acceptance evidence, including actual deployed services, two-user isolation, billing, AI budget cap, DATA integration and public HTTPS.
