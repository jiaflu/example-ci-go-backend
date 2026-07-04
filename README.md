# Example CI Go Backend

This repository is a standalone Go backend project designed to showcase a complete CI pipeline for a freelance DevOps portfolio.

## Problem

Backend teams often have tests, formatting, coverage, race checks, and build validation split across local machines. That makes pull requests inconsistent and production releases riskier.

## Solution

This project demonstrates a practical CI quality gate for a Go service:

- Go formatting check
- `go vet` static analysis
- Unit tests
- Race detector test run
- Coverage profile generation
- Coverage threshold enforcement
- Binary build validation
- CI artifact upload

## Pipeline

```text
Pull Request
   |
checkout
   |
setup Go
   |
download modules
   |
format check
   |
go vet
   |
unit tests with coverage
   |
coverage threshold
   |
race detector
   |
build backend binary
   |
upload reports
```

## Stack

- Go
- GitHub Actions
- Standard library HTTP server
- No runtime third-party dependencies

## Run Locally

```bash
go test ./...
go test -race ./...
go run ./cmd/server
```

Open:

- `http://localhost:8080/healthz`
- `http://localhost:8080/readyz`
- `http://localhost:8080/orders/summary`

## Run CI Checks Locally

```bash
./scripts/ci-local.sh
```

## GitHub Actions

The workflow is defined in:

- `.github/workflows/ci.yml`

It runs on:

- pull requests
- pushes to `main`

## Client Value

This repository demonstrates that I can set up a maintainable CI workflow for a Go backend, including fast feedback, test coverage, build verification, and reviewable CI artifacts.

