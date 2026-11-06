# Team collaboration — public boundaries

The owner coordinates engineering with multiple assistants, including ChatGPT and Muse, using a **private Drive workspace**. The private mailboxes contain operational tasks, personal data, infrastructure and unreleased model details; they are not published here.

## Public workflow

- GitHub issues and PRs for externally shareable Academy, Captain and DATA Studio development.
- Separate maintainers and review gates for education content, model releases and inference.
- Versioned schemas and reproducible code/test receipts.
- Explicit privacy reviews before any new public integration.
- Clear task ownership to avoid simultaneous editing of the same code.
- No live private Google/Colab infrastructure or credentials in public workflows.

## Release definitions

**Implemented** means source exists. **Tested** means tests ran and evidence is available. **Deployed** means its live URL was verified. **Available to learners** means a signed-out user can access it and the service works. These statuses are different.

## Roles

Academy team: original lessons, catalog review and accessible UI.
Captain runtime team: model verification, safe tool execution and inference.
DATA Studio team: evaluation and developer tooling with independently installable public code.
Owner: final authority over branding, private intellectual property and publication.
