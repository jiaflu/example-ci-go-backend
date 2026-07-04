package main

import "testing"

func TestParseCoverage(t *testing.T) {
	profile := `mode: atomic
github.com/jiaflu/example-ci-go-backend/internal/orders/service.go:10.39,12.19 1 1
github.com/jiaflu/example-ci-go-backend/internal/orders/service.go:12.19,14.3 1 0
github.com/jiaflu/example-ci-go-backend/internal/orders/service.go:16.2,20.3 4 2
`

	covered, total := parseCoverage(profile)
	if covered != 5 {
		t.Fatalf("expected 5 covered statements, got %d", covered)
	}

	if total != 6 {
		t.Fatalf("expected 6 total statements, got %d", total)
	}
}
