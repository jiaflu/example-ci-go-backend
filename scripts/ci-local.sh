#!/usr/bin/env sh
set -eu

mkdir -p .cache/go-build
mkdir -p .cache/go-mod
export GOCACHE="$PWD/.cache/go-build"
export GOMODCACHE="$PWD/.cache/go-mod"

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
./scripts/coverage-report.sh reports/coverage.out 80

echo "enforcing coverage threshold"
go run ./scripts/coverage-check.go reports/coverage.out 80

echo "running race detector"
go test -race ./...

echo "building binary"
go build -o dist/example-ci-go-backend ./cmd/server
