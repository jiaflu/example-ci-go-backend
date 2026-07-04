# Upwork Service Package

## Service Title

I will set up GitHub Actions CI for your Go backend

## Short Description

I will add a reliable CI pipeline to your Go backend so every pull request is checked before merge. The pipeline can include formatting, `go vet`, unit tests, coverage reporting, coverage threshold enforcement, race detection, binary build validation, Docker image build validation, and a short handoff document.

## Target Buyer

This service is for teams or founders who have a Go backend and want safer pull requests, more reliable releases, and a CI workflow their team can maintain.

Best fit:

- Go API/backend projects
- Small SaaS teams
- Startup engineering teams
- Solo founders with a Go service
- Teams currently running tests manually

Not a good fit:

- Projects without repository access
- Projects that need full deployment/CD setup in the same scope
- Projects with broken application code that cannot currently compile

## Core Outcome

After delivery, the client should have a GitHub Actions CI workflow that runs automatically on pull requests and `main` branch pushes.

The workflow will show clear pass/fail results for:

- Code formatting
- Static analysis
- Tests
- Coverage
- Race detection
- Build validation
- Docker image build validation

## Packages

### Basic

Goal: Add a clean CI quality gate for a small Go backend.

Deliverables:

- GitHub Actions workflow
- `gofmt` check
- `go vet`
- `go test ./...`
- Build validation
- Basic README notes

Suggested price:

```text
$150-$300
```

Suggested delivery:

```text
2 days
```

### Standard

Goal: Add CI with coverage and local reproducibility.

Deliverables:

- Everything in Basic
- Coverage profile generation
- Coverage threshold enforcement
- Race detector run
- Dockerfile review or Docker image build validation
- Local CI script
- Client handoff document

Suggested price:

```text
$300-$600
```

Suggested delivery:

```text
3-4 days
```

### Premium

Goal: Add production-grade CI cleanup for a Go backend team.

Deliverables:

- Everything in Standard
- CI runtime optimization
- Go module cache setup
- Multi-architecture Docker image build with Buildx
- Registry push setup, if registry credentials are available
- Branch protection recommendations
- Test/coverage gap notes
- Follow-up improvement plan

Suggested price:

```text
$600-$1,200
```

Suggested delivery:

```text
5-7 days
```

## Scope Boundaries

Included:

- GitHub Actions CI setup
- Go test/build pipeline
- Coverage threshold
- Dockerfile and image build validation
- Local CI script
- Documentation

Not included unless agreed:

- Production deployment
- Kubernetes setup
- Terraform infrastructure
- Registry setup or paid container registry configuration
- Database migration redesign
- Large test suite rewrite
- Fixing unrelated application bugs

## Demo Repository

This repository demonstrates the Standard package.

It includes:

- A Go HTTP backend
- Health and readiness endpoints
- Business logic with tests
- GitHub Actions CI
- Local CI script
- Coverage threshold enforcement
- Build artifact generation
- Multi-architecture Docker image build validation
- OCI image artifact export when no registry is available
- Disabled registry push placeholder
