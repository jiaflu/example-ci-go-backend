# Upwork Delivery Checklist

Use this checklist before marking the project complete.

## Repository Changes

- [ ] CI workflow added under `.github/workflows/`
- [ ] Workflow runs on pull requests
- [ ] Workflow runs on `main` branch pushes
- [ ] Go version is pinned
- [ ] Module cache is enabled where appropriate

## CI Checks

- [ ] `gofmt` check runs
- [ ] `go vet ./...` runs
- [ ] `go test ./...` runs
- [ ] Coverage report is generated
- [ ] Coverage threshold is enforced, if agreed
- [ ] Race detector runs, if agreed
- [ ] Binary or Docker build validation runs, if agreed

## Documentation

- [ ] README explains the CI workflow
- [ ] Local CI command is documented
- [ ] Branch protection recommendation is provided
- [ ] Known limitations are documented
- [ ] Follow-up improvements are listed

## Final Message to Client

Hi,

The CI setup is complete.

Delivered:

- GitHub Actions CI workflow
- Formatting check
- Static analysis
- Tests
- Coverage reporting
- Build validation
- Handoff documentation

Recommended next step:

Enable branch protection so pull requests cannot merge unless the CI workflow passes.

