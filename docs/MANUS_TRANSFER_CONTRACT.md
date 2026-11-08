# Manus Transfer Contract — NeVeRHooD Academy
Revision: 2026-11-08. Engineering status: **local-static/UI preview**, not a production application.

## Nonnegotiable product brief
- Business + application: NeVeRHooD Academy.
- Founder title: AI Multimodal Engineer & Software Engineer.
- Main experience: responsive chat-first interface with hamburger sidebar, lessons/projects/history and a sandboxed Academy lesson browser.
- Free course material and evaluation lab remain accessible without login.
- Paid Academy Plus: $10/month proposed, with **bounded provider-managed AI** when a backend and billing model are connected; never claim unlimited messages or compute.
- Captain AEGIS: separate Hugging Face showcase at https://huggingface.co/Henrybarton/captainaegis-flagship; never claim 24/7 GPU uptime.
- Aegis Data Frameworks: separate private evaluation and backend technology, not an invitation to import the whole research repo or expose its endpoints/credentials.
- Manus is the final **intended** hosting/integration platform once the offline deliverables and operator account access are ready; do not spend iterative builder credits on wholesale re-creation.

## Import bundle
- `MANUS_HANDOFF.json`: machine manifest with entries, contract boundaries and validation commands.
- `site/index.html`: generated course index (run `python scripts/build.py`).
- `site/workspace/index.html`: main UI. Includes mobile menu, tab navigation, chat draft composer (disabled until authenticated backend), same-origin lesson preview, model/profile links.
- `site/membership/index.html`: prelaunch pricing.
- `site/console/index.html`: verified status evidence.
- `site/track/foundations/index.html`, `site/learn/*` and `site/lab/evaluation/index.html`: original course/lesson materials.
- `services/academy-api/`: portable Go API; **not** yet deployed or authenticated for real learner accounts.
- `docs/ACADEMY_SUBSCRIPTION_AUTH_AI_ARCHITECTURE.md`: security and billing design.
- `docs/ACADEMY_PRODUCT_AND_REVENUE_LANES_20261108.md`: canonical owner product direction.
- `tests/`: static and behavioral regression gates.
- Main source remains in GitHub; the ZIP is an export snapshot, not an independent source-of-truth.

## Backend integration adapter to implement before opening chat
The UI's disabled send button must remain disabled until all gates pass.
- `GET /api/v1/health`: returned by the Academy Go API (health != Captain availability).
- `GET /api/v1/session`: authenticated user profile, subscription entitlement, remaining model budget and verified permissions. 401 if no active session.
- `POST /api/v1/tutor/chat`: takes a bounded question and chosen **server-allowlisted** provider route; backend obtains the provider key from a protected secret store; never expose to browser. Return grounded `answer`, `model`, `source`, `usage` or a typed refusal/unavailable status. This public contract must be redesigned from the existing **single shared Academy bearer key** before multi-tenant go-live.
- `GET/PUT /api/v1/progress/:course`: row scoped to verified user ID; optimistic version control to prevent accidental overwrites.
- `POST /api/v1/billing/checkout`: authenticated and CSRF-safe creation of price from server configured Stripe price ID, not client-supplied amount.
- `POST /api/v1/billing/webhook`: raw-body signature check, event-id dedup, entitlement recomputation, replay/reconciliation.
- `GET /api/v1/models`: public-safe model catalog and **truthful deployment state**. `not_deployed` means do not offer a working model chat.
Every request: request ID, trace classification, size/token caps, per-user rate/quota, timeout/cancel; no non-allowlisted external URLs.

## Provider-managed AI billing
- Manus or alternative provider must show exactly: supported models, monthly integration credits, per-call token pricing, AI credit threshold behavior and API quotas.
- $10 price is a **product hypothesis**, not sufficient evidence of profitability. Compute `member_revenue - payment_fees - hosting - mean_member_model_cost - support`.
- Must cap AI calls at both **member account** and **operator organization** levels.
- No automatic passing a Google identity token as a Colab credential. Google login is authentication; Google API scope authorization is a separate explicit OAuth step. Consumer free Colab is not an appropriate underlying paid API server.
- Provider-owned credentials stay in host-managed encrypted secrets. Do not put raw tokens in front-end JavaScript, localStorage, Base44 entity records, GitHub or Drive manifests.

## UI acceptance and transfer verification
1. Source review: no malicious imports, hidden outside calls, Google tokens, raw keys or fake model responses.
2. `python -m unittest discover -s tests` passes.
3. `node --test tests/evaluation_lab.test.cjs` passes.
4. `python scripts/build.py` passes.
5. Manual browser: index -> workspace -> membership -> lessons -> evaluation; mobile hamburger opens/closes; Study/Projects/Models tabs show correct content; iframe never loads arbitrary submitted URLs.
6. Keyboard: Escape closes drawer, focus indicators, navigation reachable without a mouse; desktop and narrow phone widths.
7. With no configured AI/auth, chat must visibly be offline, checkout unavailable; do not fabricate identity or AI output.
8. With host integration only after auth/entitlements, validate actual model response and usage receipts; not an offline demo.
9. Validate live custom/subdomain HTTPS signed-out, signed-in role separation, production secrets, 403/429/502, retry/recovery; only then move deployment score.
10. Owner authorizes legal terms, Stripe activation, domain/plan payment and launch; retain explicit rollback export from Manus.

## Site map
```
/                           Landing / course directory
/workspace/                 Main chat + sidebar + lesson browser
/membership/                Proposed Free / Academy Plus
/console/                   Honest engineering status
/track/foundations/         Curriculum
/learn/ai-agents-101/
/learn/protocols-101/
/learn/evaluation-101/
/lab/evaluation/             Local calculation practice
/captain/                    Interactive Aegis 12-facet architecture
/about/                      Portfolio
```

## Success definition
We are **not** at 100% when the static UI compiles. The Manus handoff is complete as an artifact when packaging, links, browser tests and build pass. Product launch readiness requires live signed-in subscriptions and independently bounded/verified AI usage. This distinction protects the owner from wasting builder credits.
