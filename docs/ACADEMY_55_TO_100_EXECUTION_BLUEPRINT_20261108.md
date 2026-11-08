# NeVeRHooD Academy — 55 → 100 execution blueprint
**Version:** 2026-11-08 | **Classification:** owner-facing executable roadmap; no claim of live deployment | **Engineering lead:** ChatGPT | **Product owner:** Henry Barton

## Objective and exclusions
Ship a secure, source-controlled, low-cost **NeVeRHooD Academy** that supports real accounts, learner records, paid $10/month Academy Plus subscriptions, bounded provider-managed AI tutoring, high-quality learning/evaluation, and a polished chat-first UI; then transfer only a verified, documented implementation to **Manus** for final provider integration/hosting. Showcase independently tested **Aegis/Captain AEGIS** on Hugging Face without requiring a hosted Captain GPU for Academy membership. Keep **Aegis Data Frameworks** as the intended evidence/evaluation backend; only explicitly sanitized, permission-checked services can cross its private boundary.

Do **not** scope in the giant Muse AEGIS brain/training project, the later enterprise proprietary application, unrestricted browser/computer agents, general student Google Colab token provisioning, unlimited hosted AI tokens, dedicated GPU hosting, assumed accreditation, or 24/7 promises based on free Colab.

## Score mechanics
This retains the existing owner 100-point weighted **PUBLIC PRODUCT-LAUNCH** rubric. 55 points already earned, 45 points require new evidence. Scores are milestone estimates, not proof of intelligence, a perfect security rating, or production SLA. A gate earns points only when there are a passing test receipt, source revision SHA, reproduced operation, and reviewable result. A feature cannot be counted twice. A passing mock doesn't count as production integration. The final 100 requires a live provider-verified URL, identity, test-mode/live billing checks, bounded real AI usage and recovery; work done offline raises component readiness but cannot certify hosting.

| Category | Baseline | Remaining | Target |
| --- | ---: | ---: | ---: |
| Academy product and learning | 16/25 | +9 | 25/25 |
| Captain model and showcase | 20/25 | +5 | 25/25 |
| Aegis Data Frameworks / Studio | 3/15 | +12 | 15/15 |
| Professional/public business | 6/15 | +9 | 15/15 |
| Native/API integration | 9/10 | +1 | 10/10 |
| Production, paid accounts and hosting | 1/10 | +9 | 10/10 |
| **Total** | **55/100** | **+45** | **100/100** |

### A. Academy learning: +9
- **A1 (+2) UI behavior and accessibility:** keyboard tab/escape/focus trapping in mobile drawer, responsive ~320–1440px and screen reader audit, browser panel safety, reduced motion, contrast and low-vision zoom tests. Genuine browser interaction checks on desktop and Android viewport; zero broken links, console errors and inaccessible controls.
- **A2 (+3) original curriculum:** expand beyond existing three lessons/15 questions with a complete original end-to-end AI software engineering pathway. Every lesson has outcomes, exercises, formative assessment, solution/rubric, clear non-accreditation, verified external sources and legal source attributions; coverage matrix and reproducible grader tests.
- **A3 (+2) learner workflow:** enrolled courses, draft/saved progress, resume on second device, versioned quiz attempts, completion rules and progress export. Requires real user authentication; two-account cross-tenant tests.
- **A4 (+2) workspace product experience:** multi-conversation UX, project/lesson context, searchable history, model-picker states and truthful offline/error handling; chat only enabled on authenticated, quota-checked gateway. Browser panel may render safe Academy-hosted lessons; arbitrary external web browsing requires separate sandboxed browser service.

### B. Captain/Aegis showcase: +5
- **B1 (+2) provenance and licensing gate:** original architecture/config/tokenizer/checkpoint hashes, attribution, research restrictions, rights to public distribution and third-party dependencies. Do not promote until verified.
- **B2 (+2) unbiased performance and safety evidence:** frozen prompt sets with held-out questions, refusal/unsafe behavior, multilingual or code if claimed, speed/p95, VRAM peak, reproducibility and failure receipts; separate model capability measurements from Academy platform quality.
- **B3 (+1) public model presentation:** accurate Hugging Face card, model-version/runtime compatibility, example usage, non-24/7 inference disclaimer, direct model card link, architecture explorer and optional known-working user-funded provider demo. **No requirement to rent Captain GPUs for the $10 school.**

### C. Aegis Data Frameworks: +12
- **C1 (+3) isolate/sanitize shared edition:** inventory private repo, commit provenance and OSS/license safety, redact all personal cloud files/account IDs, private data, keys, worker topology and unapproved model rights; create minimal licensed public-safe module/API contract rather than exposing the full operator.
- **C2 (+4) data-backed learner evaluation adapter:** versioned schemas for exercise runs, eval metrics, hashed artifacts/ProofReceipts, bounded jobs, idempotent submission and read-only learner projections. Secure operator bridge with service identity and per-tenant isolation; no direct student access to worker endpoints, training commands or arbitrary filesystem.
- **C3 (+3) evidence and correctness gates:** prove import → evaluate → compare → gate → immutable receipt → replay; hash verification, adverse/forged receipt refusal, contaminated-split rejection, crash/restart, retries/cancellation, event identity and negative tests.
- **C4 (+2) usable evidence cockpit:** 90-row audit/scoreboard filtered by status/scope (derived from real acceptance registry; not invented), links to source and receipts, honest unknown/refused labels, no operator secrets.

### D. Public identity, brand and customer trust: +9
- **D1 (+2) brand system:** consistent business/app spelling NeVeRHooD Academy, founder title *AI Multimodal Engineer & Software Engineer*, icon/wordmark palette, typography and accessibility; legal business representation distinct from trademarks.
- **D2 (+2) portfolio and case studies:** public Academy tutorial, Captain real GPU E2E (with bounded proof receipts and limitations), DATA Studio restricted/sanitized capabilities, reproducible screenshots/videos and source versions.
- **D3 (+2) customer contact and QR:** verify real external HTTPS app link, professional contact method, signed-out testing from a separate phone, printable QR to owned stable landing destination, Hugging Face model link; omit GitHub from printed card if owner chooses. Omitting link does **not** make a public repository private.
- **D4 (+3) customer/legal/release trust:** accurate pricing and AI credit policy, privacy/retention/deletion, consent, Terms, cancellation/refunds/auto-renewal disclosures, third-party course credentials/copyright, incident/contact/support policy; legal/tax review and account-owner acceptance required before collecting paid subscriptions.

### E. APIs/integration: +1
- **E1 (+1) complete evidence-bearing golden thread:** published immutable OpenAPI/Protobuf schemas and smoke test covering real frontend → auth → bounded provider AI → learning/evaluation → evidence/receipt → learner progress. Request, user, run and revision IDs remain correlated and no fake fallback answers. End-to-end pass and forced 401/403/429/502 tests.

### F. Real production, payment and hosted AI: +9
- **F1 (+2) identity and tenant boundary:** Google sign-in via verified provider (Firebase or Manus-supported hosted auth), backend JWT verification including issuer/audience/expiry/signature, per-UID isolation, logout/revocation strategy, account deletion and role controls; no shared public Academy bearer token. Test signed-out, two users, forged/expired tokens, forbidden admin, lost devices.
- **F2 (+2) subscription billing:** $10 product/price, server-controlled Checkout, webhook raw-body signature/idempotency/out-of-order recovery, subscription projection/reconciliation, cancellation/dunning/grace, billing portal and test-mode lifecycle. User signs up and cancels in test environment before any live charge. Real merchant/KYC, tax and policy activation remain account-owner responsibilities.
- **F3 (+2) actual managed AI:** one proven Manus-compatible/built-in platform provider in a real hosted flow, server-only secrets, per-user rate/token/credit caps, global org spend shutdown, honest 429/503, prompt/content limits, support for cancellation and usage receipt. Verify real answers, token accounting and no owner-budget runaway; unsupported models not shown as usable.
- **F4 (+3) hosted deployment and operations:** host app at verified public HTTPS URL, import/export and rollback tested, database migrations/backup/restore, health/readiness, logs/alerts, secrets rotation, dependency/security and accessibility tests, device/browser matrix, staging→production deployment and postdeploy smoke tests. Respect provider plan and commercial terms; no free Colab production web server.

## Implementation plan — execution order (parallel lanes where safe)
**Phase 0 — freeze ownership and QA baseline (already largely done).** Code of record GitHub `tooljhlb/NeVeRhOOd-Academy`, current source UI under `site/workspace`, docs and first Manus handoff ZIP; operator-owned private DATA repo remains separate. Do not mutate active Muse/worker Colab/CLI state. Keep dedicated Academy checkout and record SHA, timestamps, assets, test receipts.

**Phase 1 — complete portable Academy UI (A1,A4 partly; D1).** Fix interactive menu focus management, real sidebar behavior, tab keyboard control, full responsive chat UX, multi-conversation state containers and API client abstraction. Use mockable protocol fixtures but do not masquerade mock inference as live model. Create tests in Playwright/headless Chrome and static source. Deliver mobile/desktop screenshots.

**Phase 2 — build backend foundation (F1 prototype, A3).** Split current Go tutor shared-token gateway into versioned `/api/v1` router with auth middleware and typed errors; relational tables `users(uid PK)`, `plans`, `subscriptions`, `courses`, `enrollments`, `quiz_attempts`, `progress`, `chat_sessions`, `messages`, `usage_ledger`, `event_inbox`, `audit_receipts`. Locally test with SQLite migrations and seeded accounts; production use managed Postgres or host-native equivalent, selected before deploy. Add strict row ownership and transaction/idempotency tests. No client-controlled plan flags.

**Phase 3 — authentication (F1 production).** Implement adapter for one actual identity platform. Google provider must be separately enabled for real app domain. Backend verifies Firebase ID token or equivalent host-verified identity rather than accepting email strings. Cookie sessions need CSRF defenses, SameSite, Secure, HttpOnly; bearer routes need XSS/cross-origin protection. Test revocation/logout, stolen/expired credentials and data isolation.

**Phase 4 — subscription lifecycle (F2).** Stripe test-mode provider + bounded Checkout price IDs, webhook source verification, dedupe inbox, renewal/cancel/failure and entitlement read model, billing portal; simulate clocks and replay events. Publish pricing, refunds/terms before live payments. NEVER ask platform agent to start real money collection without valid merchant account and owner signoff.

**Phase 5 — cost-controlled managed AI (F3).** Define `ModelProvider` adapter, `model_catalog` and per-account/org budgets. Support one hosted platform provider first, model inventory from verified platform capability, no external arbitrary proxy URL. Token and credit receipts; policy-based limits, upstream timeouts, error classes and anti-abuse. For example, total operator model spend hard-limited to a configured monthly ceiling; per-member daily limits computed from observed model cost, not guessed. Provider AI charges operator account unless separate user-funded provider account is connected and proven.

**Phase 6 — Academy courses and projects (A2,A3).** Author original curriculum, practice lab exercises, server-graded quizzes, saved progress, feedback, advanced labs and effective search; verify third-party links and credential disclaimers.

**Phase 7 — DATA protected integration (C1–C4).** Never mount entire private DATA project into student-facing runtime. Extract minimal safe evaluation adapter and evidence-projection service; sandbox costly operator jobs and limit job types/TTL/resource budgets; render verified scoreboard with proof receipts.

**Phase 8 — Captain/HF model showcase (B1–B3).** Confirm all five shards and exact immutable revision with hashes, independent model eval and rights. Update model card and embed direct HF link and Tesseract architecture in Academy; if HF-backed demo not actually available, link to model downloads/instructions without a fake 'chat now'.

**Phase 9 — pre-Manus release qualification (D1–D4, E1, most F4).** Vulnerability audit, OWASP abuse cases, SAST/SCA, secret scan, browser security headers, WCAG 2.2 AA target and 200% zoom, offline/slow-connection behavior, tests across mobile/desktop, soak/recovery of local API, DB backup/restore, E2E contract suite and cost simulations. Freeze candidate tag with SHA and machine-readable release manifest. Rebuild the ZIP from exactly that tag and test its contents.

**Phase 10 — controlled Manus transfer (F4 and remaining production gates).** Owner provides Manus authenticated hosting permissions and commercial plan as needed; upload final source artifact once, not repeated prompt rewrites. Connect actual Google sign-in, Stripe and AI platform credentials through protected host settings, configure domain, activate preview/staging, compare hosted behavior to golden tests, run zero-dollar/test payments and provider quota tests. Audit postdeploy logs and backup restore. Only after explicit legal/payment approvals activate paid pricing; preserve rollback artifact.

## Hard requirements / blocking questions resolved
- Manus integration is intended but **no direct Manus connector available to this ChatGPT session**. Deployment can be made handoff-ready offline; actual Manus authorization/import is an external step until a compatible authorized tool exists.
- Base44 workspace **app creation blocked by free-tier five-app limit**; no existing app deletion or paid upgrade.
- Vercel account can list three projects but team scope detail returned 403; it is not proven ready as Academy production host.
- GitHub Actions account/billing lock previously blocked automation before jobs ran. Run isolated local tests and preserve receipts instead; do not mislabel CI unavailable as code failure.
- Google login cannot silently mint `token.json` for Colab. New scopes require separate consent. Google free Colab prohibits general web-service use and restricts remote control/distributed worker patterns; don't build a paid user GPU platform on it.
- Shared Desktop Commander is also used by separate active Aegis/Muse workers. Use isolated Academy workspace and do not touch their credentials, kernels, Colab CLI account states, processes or job queues.
- Paid Academy students will use platform-managed AI under operator metered credits. A customer's $10 fee is not itself a provider credit allowance. Google/Firebase data access and hosted models have separate usage charges.

## Test matrix / definition of done
| Gate | Required measured evidence |
| --- | --- |
| Static source | Python/Node suites + source build + file/link/license scans, zero fails |
| Browser | Small-phone/tablet/desktop screen + keyboard/zoom/screenreader target + real interaction recording |
| Identity | JWT signing/audience/expiry/rotation/2-user isolation, 401 and 403 paths |
| Database | Migrations, concurrency, idempotent progress updates, backup/restore and privacy deletion |
| Billing | Stripe test-mode success, cancel, webhook duplicate/out-of-order, dispute/grace and no duplicate entitlements |
| Tutor AI | Actual model response, per-user cap, org spend cap, errors 429/502/503, model/provenance ID |
| Data | Hash-bound evaluations, forged receipt rejection, worker isolation, approved read-only projections |
| Model | Hash/rights/compatibility, baseline/held-out measurements; HF public access to artifacts |
| Performance | Declared traffic scenario with p50/p95 latency, error ratio, throughput, DB and provider limits; no invented SLA |
| Security | Server authorization, CSRF/XSS/SSRF, secrets scanning, logging privacy, rate limiting and recovery |
| Hosting | Verified HTTPS, signed-out access, signed-in role, provider payment/inference, rollback and ops alerts |
| Customer trust | Pricing/terms/privacy/renewal/cancel/refunds, non-accreditation, working QR and support contact |
| Evidence | Tagged Git revision + receipt JSON + exact reproduction command and owner-visible links |

## Suggested evidence-driven milestone sequence (NOT pre-awarded points)
- **65/100**: +10 across Academy interaction/learning and DATA adapter foundations, only after corresponding A/C gates.
- **75/100**: +10 from secured DATA and portfolio/brand deliverables.
- **85/100**: +10 from Captain rights/eval, real multi-user accounts and product UX; these require actual receipts.
- **95/100**: +10 from real hosted AI/paid-subscription test and mature product/ops.
- **100/100**: final +5 from full hosted cross-service golden-thread, compliance/release and recovery criteria.
Actual point awards must follow the fixed category ledger, not these convenient round checkpoints; stage boundaries are for planning only.

## Team-of-agents division without conflicting workers
- **Lead/implementation:** ChatGPT writes GitHub source, tests, architectures, QA, receipts; uses isolated workspace and provider connectors only with specific permissions.
- **Aegis model workers:** existing separate agents continue model research and Colab work; no worker reassignment without explicit coordination.
- **Manus:** planned **last-mile host/integration platform**, not repetitive design engine; import and configure once offline acceptance passes.
- **Owner:** chooses legally binding company details, payment merchant/KYC/tax, consent to account OAuth permissions and scopes, spending budgets/domains, final commercial launch and terms. The engineering lead can prepare all technical settings and test fixtures but cannot silently accept external terms or purchase services.

## Definition of 100
100/100 is earned when:
1. A signed-out visitor opens a live secure Academy page;
2. a new user signs in with Google and receives isolated learner state;
3. the user completes and resumes an original course;
4. a test subscriber purchases/cancels $10 plan with verified server-side entitlement, and live-mode readiness is authorized;
5. a paid eligible user gets a real capped managed-AI tutor answer, with accurate receipt and organization spend guard;
6. model showcase links to valid HF card and does not claim unsupported public GPU;
7. Aegis Data Frameworks approved evaluations show verifiable receipts without exposing private operator;
8. security, accessibility, backup/recovery, logging, deployment rollback and disclosure checks pass;
9. business/QR/contact links work from another device; and
10. every milestone has a source SHA, tests and owner-visible evidence.

**Do not declare 100 on the strength of a screenshot, builder's ready status, a proposed API, a ZIP handoff, or a past GPU-only smoke test.**
