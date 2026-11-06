# NeVeRhOOd Academy

**An independent, open educational directory for programming, AI, systems engineering, and technical certification pathways.**

Created by NeVeRhOOd / Henry Barton. The Academy connects learners to original providers and provides a public, interactive **Captain AEGIS 12-facet Tesseract explorer**.

## Explore

- [Academy course directory](site/index.html) — searchable pilot catalogue.
- [Captain AEGIS architectural explorer](site/captain/index.html) — twelve interactive operational and learning facets.
- [Course source data](data/courses.csv) and [catalog builder](scripts/build.py).
- [Publication boundaries](docs/RELEASE.md).

## Development

```sh
python scripts/build.py
python -m http.server --directory site 8000
```

Open http://localhost:8000 to browse the site. A GitHub Pages workflow is included. The Pages setting may need enabling in GitHub repository Settings → Pages.

## Current capabilities and limits

The directory starts with four sample educational entries. A larger archive of course links is being cleaned, deduplicated, verified, and classified; **the full archive is not yet published**. These links point to third-party training providers. The Academy is **not accredited by, partnered with, or endorsed by** those organizations.

The Captain Tesseract is a working interactive **architectural explainer**, not a live neural-network performance monitor. Captain model tutoring and integrated chat are planned but not deployed here. Do not enter sensitive information or take exams through an AI agent without provider authorization. The Academy will prioritize learning support, original practice questions, and explanations—not impersonation, credential claims, or unattended completion of protected assessments.

## Privacy and licensing

This repository intentionally excludes private Colab infrastructure, personal cloud accounts, internal workers, credentials, unpublished checkpoints, and proprietary model-generation code. Course URLs remain attributable to their respective providers.

Source-code license: see `LICENSE`.
