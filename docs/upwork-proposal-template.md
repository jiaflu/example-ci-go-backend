# Upwork Proposal Template

## Short Proposal

Hi,

I can set up a clean GitHub Actions CI pipeline for your Go backend.

For this project, I would add:

- `gofmt` check
- `go vet`
- unit tests
- coverage reporting
- optional coverage threshold
- optional race detector
- build validation
- Docker image build validation
- optional multi-architecture image build
- short handoff documentation

I have a demo repo showing the same CI structure here:

```text
https://github.com/jiaflu/example-ci-go-backend
```

A quick question: do you want CI to only validate pull requests, or should it also run on every push to `main`?

Best,
Jiafu

## More Detailed Proposal

Hi,

I reviewed your request and this is a good fit for a focused CI setup.

My approach would be:

1. Review your current Go project structure
2. Confirm the Go version and test commands
3. Add a GitHub Actions workflow
4. Add formatting and static checks
5. Run tests with coverage
6. Add build validation
7. Add Docker image build validation if the app is containerized
8. Document how your team should use and maintain the workflow

If useful, I can also add a local script so developers can run the same checks before pushing.

Relevant demo:

```text
https://github.com/jiaflu/example-ci-go-backend
```

This demo includes a Go backend with CI for formatting, vet, tests, coverage threshold, race detector, binary build, multi-architecture Docker image build, and CI artifacts.

Questions before starting:

- Which Go version should CI use?
- Should coverage have a minimum threshold?
- Do you deploy with Docker, and which architectures do you need?
- Do you already have an image registry, or should registry push stay disabled for now?
- Do you want branch protection recommendations after CI is ready?

Best,
Jiafu

## Fixed-Price Message After Client Replies

Thanks. Based on that scope, I can deliver:

- GitHub Actions CI workflow
- Go format/static/test/build checks
- Docker image build validation
- Coverage setup
- Local CI command
- Handoff notes

I can complete this in:

```text
2-4 days
```

Before I start, I will need repository access and confirmation of the target Go version.
