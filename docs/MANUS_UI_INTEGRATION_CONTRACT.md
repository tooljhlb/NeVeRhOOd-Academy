# NeVeRhOOd UI Integration Contract (Manus or any frontend)

**Version:** 0.1.0 · **Owner:** NeVeRhOOd · **Updated:** 2026-11-06

This is the architecture-neutral integration specification for a public UI built with Manus (or another web frontend). It is **not** a claim that each backend endpoint is deployed.

## Public application sections

1. **Academy**: searchable courses and providers, original lessons, quiz exercises, learning tracks, assignments, progress, and clearly labeled *non-accredited* Academy completion certificates.
2. **Captain AEGIS**: live streamed text chat **after model inference deployment**, code/explanation support, token usage/error status, citations when retrieval is enabled; optional safe tool activity.
3. **AEGIS DATA Studio**: inspect source, run schema checks/evaluations, browse evidence receipts, compare runs and validated benchmark outputs. Only public-safe independent backend adapters.
4. **Explore**: accessible 12-facet Tesseract-inspired architecture artwork as a secondary visualization, not live cognition telemetrics.
5. **Portfolio**: technical documentation, model lineage, benchmarks, release receipts and professional links.

## Distinct deployable services (proposed)

- `web`: Manus-authored frontend. Hosted using its approved public deployment. Never embed private API credentials.
- `academy-api`: published catalog + original courses + learner account state. Can start static with `site/courses.json`; record schema `courses/schema.json`.
- `captain-api`: independently hosted inference process on compatible GPU. It loads **only** the complete validated HF model checkpoint, not the owner's private Google/Colab fabric.
- `data-studio-api`: evaluation, file analysis and verified reports; no private executor exposed to internet.
- `event-gateway`: typed job events/receipts, acknowledgments, retry behavior; server-side authorized operations only.

## Proposed, not yet deployed API shapes

```http
GET /api/v1/courses?query=&category=&verified_only=true
GET /api/v1/courses/{id}
GET /api/v1/learning-paths   # IMPLEMENTED in the read-only Go service (not hosted)
GET /api/v1/health
GET /api/v1/models
POST /api/v1/chat/completions
GET /api/v1/jobs/{job_id}
GET /api/v1/jobs/{job_id}/events
GET /api/v1/receipts/{receipt_id}
```

The public API contract uses JSON at web boundaries for interoperability; internal services may use versioned Protobuf for control/messages and FlatBuffers for read-heavy hot paths. Avoid Base64 encoding of whole large model weights or media files as a general transport solution; use resumable hashed artifacts/object storage.

## Chat transport

- OpenAI-shaped `messages`, `model`, `stream`, and bounded generation options.
- True incremental SSE for `stream=true`, finishing with `[DONE]`; never fabricate model responses from UI placeholder text.
- Server validates identity, message limit, allowed tools, maximum tool rounds, model identity and per-user rate limits.
- Structured assistant tool proposals are **never** executed merely because a model emits JSON. Only server-side allowlisted Wand interfaces can act.

## Credential and learner safeguards

- External university certificates/credits are awarded solely by the provider. The Academy cannot award a PhD.
- Original Academy completion certificates must be clearly non-accredited and attributable to a completed assessed course. Require signed course progress/evidence before issuing.
- Don't perform restricted third-party tests on behalf of students.
- Never request private Google account credentials through public UI.
- Separate customer-facing login from owner/admin cloud credentials; no browser-accessible infrastructure control-plane tokens.

## Workflow/mailbox (private operator only)

The authoritative owner coordination area is **private** at:
`G:\My Drive\aegis-env-master-local\NeVeRhOOd Academy\03_MAILBOX`.

Suggested future event envelope (schema draft):
```json
{
  "schema_version": 1,
  "event_id": "uuid",
  "task_id": "uuid",
  "attempt": 1,
  "idempotency_key": "unique-within-project",
  "event_type": "task.started",
  "resource_ref": "receipt-or-job-id",
  "occurred_at": "2026-11-06T00:00:00Z",
  "actor_role": "system",
  "sha256": "64-hex-digest"
}
```

`task.started` and `task.completed` must reflect genuine server-confirmed events. The frontend may subscribe via SSE/WebSocket; it does not read the private Drive mailbox directly. Enforce audit logs, cancellation, idempotency, retry budgets and redaction.

## Staged UI states

Show *Available* only for endpoints that respond successfully. Show *Preview* for static/Tesseract content, *Under development* for absent backends, and *Unavailable* with error details for failed services. A static mock of Captain chat is not a working model demo.

## Launch gates

- Academy: original lessons/quiz operational; source links correct and provider terms reviewed; accessibility responsive audit; live site URL verified.
- Captain: all five root SafeTensors shards verified, tokenizer/chat template and real CUDA inference tested; authorization, rate limits and streaming verified.
- DATA Studio: independent public-safe runtime built with sample evaluation receipts, without owner's private Google compute integration.
- Manus UI: tested against documented responses, loading/error states, signed-out visitor test and mobile/low-vision accessibility.
- Portfolio: public model card links, repo links, release measurements and honest limitations.

## Current technical blockers

As of last verified session: GitHub Pages workflow failing; Hugging Face only shard 1 of 5 confirmed in destination, full GPU inference absent; HF Gradio Space creation blocked by account hosting requirement (HTTP 402). Do not hardcode a live backend URL before it's established.

## Proprietary boundary

Do not expose owner's local Windows filesystem or private Google Drive mailboxes, custom Colab worker fabric, private source/repos, internal Cloudflare workers, model generation secrets, unpublished media packs or unreviewed personal records.

## Foundations track — implemented November 6, 2026

The public static source now includes [AI Engineering Foundations](../site/track/foundations/index.html), with three original lessons and fifteen practice questions. The independently testable Go API provides `GET /api/v1/learning-paths` returning an `ai-engineering-foundations` record with three lesson IDs and URLs, `accredited: false`, and `certificate_status: not_issued`. It does not collect learner identities or award credentials. Use these honest states in Manus.

The refreshed Manus static website ZIP is stored privately in the owner workspace under `10_RELEASES/NEVERHOOD_ACADEMY_MANUS_FOUNDATIONS_20261106.zip`; it is not a live deployment. GitHub Actions is restricted at the account level; the site builds locally.
