# Upwork Buyer Requirements

Send these questions to the client before starting.

## Access

- Do you use GitHub?
- Can you invite me to the repository?
- Should I work on a feature branch and open a pull request?

## Project

- What Go version should CI use?
- Does `go test ./...` currently pass locally?
- Does the project require private Go modules?
- Are there required environment variables for tests?
- Does the project need a database, Redis, or another service for tests?
- Is the app deployed as a Docker container?
- Do you already have a Dockerfile?

## CI Behavior

- Should CI run on pull requests?
- Should CI run on pushes to `main`?
- Should coverage have a minimum threshold?
- Should the race detector be included?
- Should CI build a binary or Docker image?
- Which image platforms do you need, such as `linux/amd64` or `linux/arm64`?
- Do you have an image registry, such as GHCR, Docker Hub, ECR, GCR, or ACR?
- Should registry push be enabled now or left as a placeholder?

## Delivery

- Do you want branch protection recommendations?
- Do you want a short handoff document?
- Should I add a local CI script for developers?

## Red Flags

Clarify scope before accepting if:

- Existing tests do not pass
- The app does not compile
- The repo requires secrets the client cannot provide
- The client expects deployment/CD in the same small CI scope
- The client wants unrelated application bugs fixed under the CI task
