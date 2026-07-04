# CI Demo Cases

This repository has two CI modes:

- Real CI: `.github/workflows/ci.yml`
- Manual demo cases: `.github/workflows/ci-demo-cases.yml`

## Real CI

The real CI workflow runs on pull requests and pushes to `main`.

It should stay green because it protects the repository.

## Demo Cases

The demo workflow is manual only. It is designed for portfolio and client demos.

Open:

```text
Actions -> CI Demo Cases -> Run workflow
```

Choose one of:

- `success`: runs the normal green path
- `failure`: intentionally fails each check but continues through the rest of the pipeline
- `both`: runs both cases

## Failure Case Behavior

The failure case uses `continue-on-error: true` on each demonstration step.

That means a failed formatting check does not stop the test failure demo, coverage failure demo, Docker image failure demo, or registry failure demo from running.

Bootstrap steps such as checkout, Go setup, QEMU setup, and Buildx setup are allowed to fail normally because later steps depend on them.

At the end, the job writes a summary table and then fails intentionally so the workflow visibly shows a failed demo result.
