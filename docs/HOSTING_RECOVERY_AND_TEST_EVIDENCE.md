# Academy deployment recovery and independently verified tests

**Recorded:** 2026-11-06

## Confirmed GitHub Actions provisioning failure

Public GitHub Actions run [42311455409](https://github.com/tooljhlb/NeVeRhOOd-Academy/actions/runs/42311455409) failed before any job steps ran.

The job's GitHub check-run annotation (ID `128213623257`) states:

> The job was not started because your account is locked due to a billing issue.

The adjacent runner-image migration notice is informational, **not** this failure's cause. Re-running the same workflow without resolving the account restriction will not validate or deploy the site.

**Account owner action:** review GitHub's account billing/status and Actions restrictions, or contact GitHub Support. Do not put payment information, recovery codes, access tokens or billing secrets in this repository or in team messages. An account lock does not automatically imply that a paid plan is required; follow GitHub's account-specific resolution.

## Independent validation (not GitHub Actions)

The following checks were run locally from source corresponding to the Academy repository:
- Python catalog validation: **4 records, 4 distinct destinations**, all structurally valid (provider academic claims remain unverified).
- Python `unittest`: **6 tests passed**.
- Static site generator: **passed**, produced HTML and course JSON.
- Go API: reconstructed matching main source and six tests in an independent Go 1.23.2 environment; **6 Go tests passed**, `go vet` passed, race-detector tests passed, and executable compilation passed. This verifies source behavior under the test inputs, **not** a production hosted instance or current GitHub CI.

The Go tests check read-only behavior, health, searches/pagination, unverified-only filters, CORS allowlisting and explicit `captain-aegis: not_deployed` response.

## Temporary delivery path while Actions is locked

1. Build the reviewed static site locally: `python scripts/validate_catalog.py && python -m unittest discover -s tests -v && python scripts/build.py`.
2. Deliver **only the contents of the `site/` directory** as a static artifact to Manus or another explicitly authorized public host.
3. If Manus deploys it, check the actual HTTPS URL signed out on desktop and mobile, original quiz, links, and keyboard accessibility. Record deployment receipt and commit SHA.
4. Configure an independent Go API host when a supported runtime and public HTTPS endpoint are available; the static site **does not require the Go API** to function.
5. Do not turn a static mock chat into a claimed live Captain inference experience. The Captain model and Muse's large AEGIS model are separate.
6. Once Github account restrictions are resolved, re-enable/test GitHub Actions and Pages as desired; Manus may remain the main public UI host.

## Operator safety

This Academy work does not operate the owner's private Colab accounts, AEGIS generator, GPU jobs, or credentials. Public code and source artifacts only; no HF weight duplication in Drive.
