# NeVeRhOOd Academy

**An independent, open educational directory for programming, AI, systems engineering, and technical certification pathways.**

Created by NeVeRhOOd / Henry Barton. The Academy connects learners to original providers and provides a public, interactive **Captain AEGIS 12-facet Tesseract explorer**.

## Captain AEGIS on Hugging Face

- [Public Captain AEGIS model repository](https://huggingface.co/Henrybarton/captainaegis-flagship) — five SafeTensors shards, [real Tesla T4 inference proof](https://huggingface.co/Henrybarton/captainaegis-flagship/blob/main/evidence/captain_t4_gpu_smoke_20261106.json), and [native authenticated API/SSE proof](https://huggingface.co/Henrybarton/captainaegis-flagship/blob/main/evidence/captain_t4_native_api_20261106.json). A publicly hosted chat service is **not** yet deployed.

## NeVeRhOOd public channels

- [NeVeRhOOd Academy — YouTube](https://www.youtube.com/@NeVeRhOOd_Academy)
- [Public engineering portfolio](docs/NEVERHOOD_PUBLIC_PORTFOLIO.md)
- [Academy Go API](services/academy-api/README.md) (source; not hosted)

## Launch status

[Consolidated release tracker](docs/MASTER_RELEASE_TRACKER.md) — verified milestones for the school, Captain checkpoint, public runtime, HF chat demo and professional portfolio.\n\n## Start learning

- [AI Agents 101 — Tools, Permissions & Evidence](site/learn/ai-agents-101/index.html): original five-section lesson with interactive five-question practice quiz; independent and non-accredited.
- [AI Engineering Foundations](site/track/foundations/index.html): three original lessons, fifteen self-assessment questions.
- [Protocols 101](site/learn/protocols-101/index.html): JSON, Protobuf and FlatBuffers.
- [Model Evaluation 101](site/learn/evaluation-101/index.html): reproducibility, contamination, controls and release gates.
- [Interactive Model Evaluation Lab](site/lab/evaluation/index.html): edit confusion-matrix counts, inspect precision/recall/F1/accuracy, export a local practice receipt.
- [Academic credits and certificates policy](docs/ACADEMIC_CREDIT_AND_CERTIFICATES.md)
- [Curriculum roadmap](docs/CURRICULUM_ROADMAP.md)

## Explore

- [Manus-portable Academy study workspace](site/workspace/index.html) — mobile hamburger navigation, offline chat shell, project tabs and same-origin lesson browser; AI not connected yet.
- [Academy Plus membership preview](site/membership/index.html) — proposed $10/month with usage-limited platform AI; no live checkout or paid AI access yet.
- [Owner-approved product and revenue lanes](docs/ACADEMY_PRODUCT_AND_REVENUE_LANES_20261108.md) — Academy, Aegis Data Frameworks, Aegis model showcase and separate business services.
- [Academy Control Center](site/console/index.html) — responsive product dashboard with evidence-based status and safe optional API health check.
- [Academy course directory](site/index.html) — searchable pilot catalogue.
- [Captain AEGIS architectural explorer](site/captain/index.html) — twelve interactive operational and learning facets.
- [Course source data](data/courses.csv) and [catalog builder](scripts/build.py).
- [Publication boundaries](docs/RELEASE.md).

## Public portfolio and engineering boundaries

- [NeVeRhOOd portfolio page](site/about/index.html) — Academy, Captain AEGIS and DATA Studio
- [Program boundaries](docs/PROGRAM_BOUNDARIES.md) — Muse's separate AEGIS model and private Colab work are strictly excluded
- [Standalone Go Academy API](services/academy-api/README.md) — source available, public hosting pending; local Go tests passed
- [Open release issues](https://github.com/tooljhlb/NeVeRhOOd-Academy/issues) — deployment, academic review, Captain integration and Manus

## Development

```sh
python scripts/build.py
python -m unittest discover -s tests -v
node --test tests/evaluation_lab.test.cjs
python -m http.server --directory site 8000
```

Open http://localhost:8000 to browse the site. A GitHub Pages workflow is included. The Pages setting may need enabling in GitHub repository Settings → Pages.

## Current capabilities and limits

The directory starts with four sample educational entries. A larger archive of course links is being cleaned, deduplicated, verified, and classified; **the full archive is not yet published**. These links point to third-party training providers. The Academy is **not accredited by, partnered with, or endorsed by** those organizations.

The Captain Tesseract is a working interactive **architectural explainer**, not a live neural-network performance monitor. Captain model tutoring and integrated chat are planned but not deployed here. Do not enter sensitive information or take exams through an AI agent without provider authorization. The Academy will prioritize learning support, original practice questions, and explanations—not impersonation, credential claims, or unattended completion of protected assessments.

## Privacy and licensing

This repository intentionally excludes private Colab infrastructure, personal cloud accounts, internal workers, credentials, unpublished checkpoints, and proprietary model-generation code. Course URLs remain attributable to their respective providers.

Source-code license: see `LICENSE`.
