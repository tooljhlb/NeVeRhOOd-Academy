# NeVeRhOOd Academy — Subscription, identity and AI provider architecture

Status: **engineering design / not deployed**. Created November 7, 2026. No live billing, Google OAuth or paid AI usage is implied by this document.

## Product decision
- Name: NeVeRhOOd Academy, operator NeVeRhOOd, founder Henry Barton.
- Founder title: **AI Multimodal Engineer & Software Engineer**.
- Target launch plan: **Academy Plus, $10/month**, with free introductory learning.
- No unlimited Captain GPU inference bundled in the $10 subscription.
- Public model: https://huggingface.co/Henrybarton/captainaegis-flagship
- GitHub Academy: https://github.com/tooljhlb/NeVeRhOOd-Academy
- User-oriented links to prepare: Academy home, Learn, AI Playground, Captain Model, Pricing, Developer Profile, YouTube.
- Business-card QR: one stable *public HTTPS landing URL* eventually linking to these; until that exists, Academy GitHub URL is a valid interim destination. Do not encode access keys in QR or static HTML.

## Services and responsibilities

| Service | Responsibility | Trust/hosting boundary |
| --- | --- | --- |
| Academy frontend (GitHub owned; Base44 or alternative) | Pages, account UX, progress views, model picker and pricing | Treat all browser data as untrusted. No provider keys or unrestricted API proxy in shipped JS |
| Firebase Auth (Google provider) | Authenticate identity | Client signs in, Go backend **verifies Firebase ID JWT** against Google's issuer/audience/signature and checks expiry and revocation policy where appropriate |
| Academy Go API | Canonical learner/plan/progress authorization, quotas, model routing and audit receipts | Authorization on **every** protected request; never trust client plan flag, client subtotal or provider model selection |
| Academy DB | Users keyed to immutable Firebase UID, subscription entitlements, usage ledger, enrollment/progress, provider connection metadata | Server-side access only; row isolation per verified UID; migration/restore testing |
| Stripe Billing | Hosted checkout, subscriptions, customer portal, receipts and tax configuration | Backend creates Checkout Sessions; verify webhook signature over **raw body**, dedupe event ID and process asynchronously; Stripe API secrets never reach browser |
| BYOK model gateway | For eligible learner-owned API accounts; proxy requests safely | Host allowlist and fixed provider catalog; per-user encrypted credential storage or provider OAuth, budget/rate limits, no raw URL supplied by customer |
| Captain inference gateway | Independently hosted Captain GPU, allowlisted tools, metering | Default OFF until durable GPU hosting with authentication, rate limits, cancellation and recovery |
| Hugging Face | Model weights and evidence; optional external provider billing | HF model repository alone does **not** provide 24/7 inference. HF routed calls generally bill the token/account owner; user's own token must be used for user-paid requests |

## Subscription proposal
### Free
- Public course discovery and original foundational lessons; in-browser practice evaluation.
- Documentation and model card links, *not* guaranteed hosted AI model inference.
- Optional Google account for progress after identity stack is deployed.

### Academy Plus: $10 per month (planned, not accepting payments yet)
- Saved learner progress and private workspaces.
- Additional original curriculum, enhanced evaluation labs and project templates.
- Provider-model playground with **bring-your-own-key** (BYOK), where the owner of the provider account pays their own provider bills.
- Premium content access controlled by server-side entitlements; do not imply third-party credentials or accredited diplomas.
- Optional separately purchased Captain hosted inference only **after** unit economics, GPU availability, protections and refund policy are validated.
- No promise of unlimited AI tokens, premium provider credits or guaranteed dedicated GPU.

### Revenue and cost model
Current official Stripe US reference at time of design: Payments domestic card 2.9% + $0.30/charge; Stripe Billing pay-as-you-go 0.7% of recurring billings. Illustrative $10 transaction: $10 - $0.29 - $0.30 - $0.07 = **$9.34 before tax handling, refunds, disputes, hosting, support and API/compute**. Verify actual account pricing before enabling.
Gross plans: 25 members => $250 MRR; 100 => $1,000 MRR; 500 => $5,000 MRR. These are revenue scenarios, not projections.
Risk: a $10 account consuming >$9 of operator-paid GPU/API time is unprofitable before overhead. Quota all operator-paid inference; meter tokens and enforce spending caps.

## Authentication flow
1. Browser uses Firebase Auth Google sign-in; capture Firebase ID token (NOT Google provider access token) to call Academy Go API.
2. Go checks issuer https://securetoken.google.com/<project-id>, audience <project-id>, signature from Google's current keys, exp/iat, subject UID and tenant restrictions as applicable; reject invalid tokens.
3. Server upserts profile by Firebase UID, not client-provided email; default FREE. Email is an attribute, not the stable authorization key.
4. Browser never specifies authoritative subscription role. Go reads entitlements from server-side Stripe-backed projection.
5. Admin routes require separate role claims established server-side and explicit auditing.
6. All calls over HTTPS; CSRF-safe session design if cookies; if bearer tokens, guard XSS and avoid sensitive token storage in persistent client storage.

## Billing and entitlement flow
1. Authenticated learner presses Subscribe. Go determines allowed server-side plan/price IDs; never takes arbitrary Stripe price from browser.
2. Backend initiates Stripe Checkout and sets correlation metadata to verified UID.
3. Redirect to Stripe checkout. Frontend return page is NOT proof of payment.
4. Stripe webhook verified with configured signing secret, raw body and timestamp; event ID deduplicated atomically.
5. Persist subscription status from authoritative Stripe event or reconciliation fetch; update entitlement transactionally and audit.
6. On failed invoice/cancellation/chargeback, enforce explicit grace-period and access policy; reconcile daily with Stripe.
7. Offer billing portal for cancel/update card. Tax, refunds, terms and privacy policy required before launch.

## BYOK gateway
- Users choose from **server-configured** provider/model allowlist. Never let them submit arbitrary upstream URL/host or redirect path.
- Initial allowed providers should be independently verified for BYOK terms and direct API: examples may include Hugging Face Inference Providers and compatible API vendors; no vendor enabled merely because a plugin is listed.
- Store user-provided API keys encrypted at rest using a real server-side managed secret/key service; per-user envelope keys and audit-safe key identifiers. No Base44 entity named APIKey storing raw secrets, no localStorage persistence, no logging. Offer revoke/rotate/delete.
- Prefer provider OAuth and delegated tokens where supported over capturing personal API keys.
- Keep per-user quota/timeout, max output tokens, request size, request idempotency and cancellation; enforce no cross-user access.
- Explicitly disclose billing: customer-provided provider token => provider charges **that customer's account**. Our own proxy egress/storage still costs us; cap that too.
- First public iteration can omit saved keys entirely and let customers enter short-lived keys for session-only proxy use after safeguards are implemented.
- Provider integration is separate from Captain's model. Captain via Hugging Face model card is a download/demo artifact unless a dedicated endpoint is deployed.

## Captain hosting
- Existing real private Captain native FastAPI + Academy Go integration passed on isolated Tesla T4, but free Colab is **ephemeral**, not a paid customer GPU backend.
- For Captain-hosted plan, choose reliable isolated GPU host, cost cap and scheduling, retain no arbitrary shell tools; implement admission control, auth, tokenizer quotas, daily budgets, cancellation, retry/reconnect, rate limits, abuse monitoring and incident receipts.
- Keep native GPU workers private behind API. Public clients only use Academy gateway with scoped identity and billing checks.
- Dedicated endpoint may be priced as separate metered credits once measured (token/sec, cost/token, startup time, p95 latency, cold-start).

## Platform decisions
- Base44 desired as UI builder, but **new app create returned free-tier 5-app limit**; existing 15 apps must not be deleted. External coding agent/sandbox requires Builder entitlement. No automatic purchase.
- Lovable alternative: full-stack TypeScript AI builder, would introduce an extra implementation/runtime; avoid duplicate business logic.
- Vercel team access returned 403; no deployed Academy project verified. Vercel Hobby is non-commercial under current terms; charged subscriptions require eligible hosting.
- Manus optional as a corporate landing host later, not blocked on.
- Code of record is public Academy GitHub, private Captain runtime repo, HF model assets and owner evidence.

## Acceptance checklist
- [ ] User consent to publish billing terms and privacy policy, final legal business name, tax settings and verified payment account.
- [ ] Google sign-in works on the actual authorized domain; session rotation/logout and unauthorized 401 tested.
- [ ] Firebase project and authorized OAuth domain configured by account owner.
- [ ] Database migrations and multi-tenant isolation tests pass.
- [ ] Stripe test-mode Checkout, webhook signature, replay/dedup, failed renewals, cancellation and portal pass.
- [ ] $10 pricing page clearly distinguishes free lessons, premium features and metered AI costs.
- [ ] Provider BYOK tests ensure no secret leaks in browser/network logging, no SSRF/cross-user reads, revocation and quotas.
- [ ] Captain GPU not advertised as always-on without durable host and recovery SLA tests.
- [ ] Public domain HTTPS, accessibility, analytics opt-in, backup, refund and legal pages verified signed-out.
- [ ] Paid subscription LIVE only after operator approves account activation and all gates PASS.

## References reviewed November 7, 2026
- https://firebase.google.com/docs/auth/web/google-signin
- https://firebase.google.com/pricing
- https://stripe.com/pricing
- https://stripe.com/billing/pricing
- https://huggingface.co/docs/inference-providers/pricing
- https://vercel.com/legal/terms
