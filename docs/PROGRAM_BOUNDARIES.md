# Program boundaries — locked November 6, 2026

This document separates independent engineering programs. It exists to prevent conflating another AI worker's model training with the NeVeRhOOd public portfolio.

## Lane A: AEGIS / GeneralAegis / Muse — separate

- Muse owns the owner's long-running AEGIS large-model generation, training, evidence and Google Colab worker operations.
- AEGIS uses separately authorized private Google/Colab accounts, credentials, jobs, code and checkpoint repositories.
- This Academy development lane **must not** start, stop, redirect, inspect token files, control, or compete for Muse's cluster resources or infer that AEGIS training relates to Captain.
- The owner offered a different Colab account, but **it is not required for these current Academy tasks**. Do not read its credential folder.
- Do not place AEGIS proprietary model files, worker systems, credentials or private execution interfaces in the Academy, Captain public runtime, or DATA Studio public editions.

## Lane B: Captain AEGIS — independent small public-facing model

- Captain is its own approximately 4.206B physical-parameter research checkpoint and separately tested model-serving runtime.
- Its root checkpoint is backed by an HF bucket and private staging model repository; the owner/Muse may separately handle the safe shard transfer.
- Captain is intended for a future optional Academy tutor and code/learning help, but full inference and endpoint readiness must be verified before calling the chat live.
- A model card, binary checkpoint and running inference service are **three distinct artifacts**.
- Serving Captain must not require the owner's AEGIS/Colab compute control system.

## Lane C: NeVeRhOOd Academy — ChatGPT and Manus

- ChatGPT: the public GitHub source, original courses/quiz, catalog validation, independent Go read-only API, tests, integration contract, issue tracker, release evidence and professional portfolio.
- Manus: future public frontend styling and hosting. ChatGPT supplies stable endpoints and documented fallback/unavailable states; no competing UI rewrite.
- Drive is a **lightweight team record/mailbox workspace**, not a bulk model-weight mirror.
- YouTube: https://www.youtube.com/@NeVeRhOOd_Academy
- GitHub: https://github.com/tooljhlb/NeVeRhOOd-Academy
- Brand: NeVeRhOOd Academy, graduation/learning motif and AI engineering emphasis.

## Lane D: AEGIS DATA Studio — public-safe standalone engineering tools

- Source is the existing PRIVATE repository: https://github.com/tooljhlb/aegis-DATA-framworks
- Public edition must be an independently installable, rights-reviewed release containing approved engineering/evaluation/code-analysis modules; NOT private compute or Google account integrations.
- Connection to Captain is optional only after a working, authorized public inference endpoint exists.

## What may be shared across lanes

Only deliberately published, sanitized interface contracts, generic documented algorithms, original rights-safe educational material, and measured receipts that do not expose private infrastructure.

## Release rules

1. Never claim the large AEGIS model and Captain are the same model.
2. No reading Colab tokens or using extra Google accounts for Academy work.
3. No changing Muse's long-running training jobs.
4. No public exposure of private Drive mailboxes, credentials or infrastructure.
5. No unverified model quality, university accreditation, PhD or academic-credit claims.
6. Public endpoints require independent tests, HTTPS, scoped authentication and documentation before Manus integrates them.
