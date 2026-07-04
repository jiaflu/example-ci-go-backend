#!/usr/bin/env sh
set -eu

mkdir -p .cache/go-build
export GOCACHE="$PWD/.cache/go-build"

echo "checking gofmt"
unformatted="$(gofmt -l .)"
if [ -n "$unformatted" ]; then
  echo "$unformatted"
  echo "files need gofmt" >&2
  exit 1
fi

echo "running go vet"
go vet ./...

echo "running tests with coverage"
mkdir -p reports dist
go test ./...
go test ./internal/... -covermode=atomic -coverprofile=reports/coverage.out
go tool cover -func=reports/coverage.out

echo "enforcing coverage threshold"
go run ./scripts/coverage-check.go reports/coverage.out 80

echo "running race detector"
go test -race ./...

echo "building binary"
go build -o dist/example-ci-go-backend ./cmd/server
