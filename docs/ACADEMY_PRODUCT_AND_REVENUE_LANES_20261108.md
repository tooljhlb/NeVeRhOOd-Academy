# NeVeRHooD Academy — owner product, revenue and platform decisions
Date: 2026-11-08. Status: **approved product direction, not a claim of live deployment or accepted payments**.

## Owner-defined identities (keep separate)
- **Business name:** NeVeRHooD Academy.
- **Application name:** NeVeRHooD Academy.
- **Application backend / evaluation framework:** Aegis Data Frameworks. The existing standalone Academy Go API is a working tutor gateway; it is NOT yet confirmed that Aegis Data Frameworks fully backs the deployed Academy.
- **Model-family name:** Aegis. **Captain AEGIS** is the current separate HF showcase and proof-of-inference lane. It is not a continuously hosted paid Academy backend.
- **Founder title:** AI Multimodal Engineer & Software Engineer.
- These are owner-requested labels; existing repository name/case and separate private Muse giant-AEGIS workers are not renamed or altered by this decision.

## Revenue lane A: low-cost school
- **$0 Free**: Public learning, sample lessons, hands-on examples, public Aegis HF model card.
- **$10 per learner/month (proposed) Academy Plus**: more lessons, saved progress, experiments, model-backed tutor with **strict platform-managed AI usage caps**. Not unlimited usage.
- **Primary inference option:** Base44 built-in AI integrations when the app and plan allow them. Host/admin pays platform integration credits; users do **not** automatically pay for the AI just because they paid $10. Limit calls/user/day, model mix, output tokens and daily/monthly owner spend; gate feature when remaining credits are insufficient.
- **Secondary model option:** Manus in-app AI only after checking the specific product's API/hosting plan, acceptable use and per-request costs. User's Manus account isn't exposed as an available direct Manus connector to this ChatGPT session (searched installed tools and plugin directory); no Manus deployment claimed.
- **Optional BYOK later:** Advanced members may connect a supported provider with delegated auth or a secure server-side proxy; never use localStorage or plain-text data tables for API keys, and never imply that Base44 integration credits are billed to the learner's external AI account.
- **Captain AEGIS:** Showcase separately through https://huggingface.co/Henrybarton/captainaegis-flagship and the existing public static architecture UI. Its public HF repo is not a guaranteed free hosted model endpoint. Paid Captain hosted inference is a later business case after benchmarked GPU unit economics.
- **No live billing** until verified authentication, entitlement enforcement, terms/refunds, secure Stripe/provider connection, test-mode webhook and signed-out production QA.

## Revenue lane B: local business service work (NOT school plans)
Owner's exploratory market-facing quotations, requiring real scopes, service levels, taxes, margins and contracts before publication:
- On-site/device setup: **$125/hour** example.
- Managed per-device/per-user AI IT support: sample $150/month for 1 user; $750 for 5; $1,200 for 10; $2,500 for 25; $4,000–$5,000 for 50; custom $7,500+ for 100+. These are NOT validated prevailing market prices or guaranteed profitable quotes.
- Alternative illustrative small-business managed package: $1,000–$1,500/month for ~10 users, or $250–$500/month narrower maintenance; these cover materially different service levels, not stackable or equivalent prices.
- For one-to-one sales, use **discovery + written quote** based on device count, automation scope, workload and incident SLA; avoid promising unlimited support, dedicated GPU, unlimited AI calls, regulatory compliance, or full device management without infrastructure.
- User asked to keep this service-price lane distinct from the low-cost Academy subscription.

## Revenue lane C: later professional corporation product
- Only after revenue evidence: upgraded corporate identity, a proper owned domain, bespoke Aegis model/service, enterprise hardware/hosting, billing tiers ~ $20/month personal professional and ~$250+/month business as **targets subject to measured costs**.
- Manus is owner's preferred possible professional application host; it advertises built-in auth, Stripe and domain features. Check live cost/export/lock-in and eligibility before buying or publishing.

## Business card and QR — owner approved direction
- Print **NeVeRHooD Academy**, **AI Multimodal Engineer & Software Engineer**, services (AI agents/automation/software), and a QR to a **verified public app or landing page**, once available.
- Keep GitHub **off the physical card**. Note: the current Academy repository is publicly accessible; omitting it from a card does NOT make the repository private. Do not change visibility without reviewing the public site/source and release consequences.
- Feature Aegis model card https://huggingface.co/Henrybarton/captainaegis-flagship separately and include an app URL once a signed-out working deployment exists.
- Do NOT print a guessed Manus/Base44 subdomain or an unverified QR destination.
- Owner's white-card design handled by another model; this file is the accurate URL/content handoff.

## Platform limits actually observed
- Base44 account exposed 15 apps, but new NeVeRHooD Academy creation explicitly **failed** due to free-tier limit 5 apps; do not delete an app for quota.
- Base44's remote direct file editing is blocked (Builder required). Built-in platform AI requires integration credits; Base44 free published tier shows 100 integration credits/month and five-app limit; **not unlimited free user inference**. Builder currently shows $40/mo billed annually with 10k integration credits/mo, not evidence of current plan or custom entitlements.
- Base44 Github and Stripe advanced integrations require Builder or higher according to provider docs. No nested OAuth connection currently proven for those apps.
- Vercel access returned team scope 403 for project details; not a verified Academy host.
- Manus supports website domains, login and Stripe integration in its product documentation, but a Manus tool isn't currently available in this session. No claim we connected or deployed there.
- Connected Desktop Commander is shared with a separate actively running Aegis worker/other sessions. **Do not operate shared working directories or CLI Colab state without isolated workspace, locks and explicit coordination.** Prefer GitHub branch or new worktree for Academy; no Muse lane files should change.

## P0 operational milestones for raising readiness
1. **One public live app** with account-bound production link (not an HTML prototype); use Base44 if account/credits unblock, otherwise eligible Manus/Lovable/Vercel hosting. Verify signed-out pages, navigation, mobile.
2. **Identity & entitlement**: Google login + per-user account + server-side paid/free roles + tenant isolation; valid session and 401/403/expiry/logout tests.
3. **Stripe test-mode**: $10 recurring product, checkout, signed webhook, replay/idempotency, status reconciliation, cancel/portal; then account-owner approval to enable live payments.
4. **AI tutor as bounded platform-managed credits**: prove which platform handles model calls, exact credit charge by model, max requests/user/day/month; abort on insufficient budget. Never expose owner's provider keys or claim individual learners' credits automatically cover operator spend.
5. **Aegis Data Frameworks backend integration**: stable interface/schema/metrics and redacted public-safe evidence, not simply declaring the frameworks 'connected'.
6. **External Captain Aegis model**: linked HF model card + public accurate status, optional public demonstration only when hosted or safely queued.
7. **Business-card-ready QR**: generated from a URL verified as deployed and accessible to signed-out mobile users.
8. **Promotion criteria**: no launch score increase for plans, provider catalog features, local static source or untested integrations; score rises on evidence-backed live milestones.

## Unit economics gate
At $10/month: gross 10 subscriptions $100; 100 subscriptions $1,000; 500 subscriptions $5,000.
Subscription gross is NOT contribution margin. Model requests on Base44 use operator **integration credits**; recurring revenue doesn't create per-member platform credits automatically. Get model-specific test usage receipts, monitor credits per learner, set cost caps before selling.
Use an actual usage model (N paid learners × mean messages/learner × credits/request) plus Stripe costs, platform subscription, storage, data and support. Define a hard outage/degrade behavior when credits exhaust; no subsidized unlimited AI.

## References
- Base44 pricing https://base44.com/pricing
- Base44 integrations https://base44.com/integrations
- Manus site https://www.manus.im/features/webapp
- Manus Stripe https://www.manus.im/connectors/stripe
