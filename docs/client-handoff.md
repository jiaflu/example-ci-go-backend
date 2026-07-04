# Client Handoff

## CI Checks Included

- Formatting check with `gofmt`
- Static analysis with `go vet`
- Unit tests for handlers and business logic
- Race detector test run
- Coverage profile generation
- Minimum coverage threshold
- Backend binary build validation
- Multi-architecture Docker image build validation
- OCI image artifact export while registry push is disabled
- CI artifact upload for review

## Recommended GitHub Settings

- Require pull requests before merge
- Require the `Format, Vet, Test, Coverage, Build` status check
- Require branches to be up to date before merge
- Restrict direct pushes to `main`

## Production Extensions

- Enable image registry push after registry credentials are available
- Add integration tests with a real database
- Add vulnerability scanning
- Add CodeQL
- Add deployment workflow after CI is stable
