package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/coverage-check.go <coverage-file> <minimum-percent>")
		os.Exit(1)
	}

	coverageFile := os.Args[1]
	minimum, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid minimum coverage: %v\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(coverageFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read coverage file: %v\n", err)
		os.Exit(1)
	}

	covered, total := parseCoverage(string(data))
	if total == 0 {
		fmt.Fprintln(os.Stderr, "coverage profile has no statements")
		os.Exit(1)
	}

	percent := (float64(covered) / float64(total)) * 100
	if percent < minimum {
		fmt.Fprintf(os.Stderr, "coverage %.2f%% is below required %.2f%%\n", percent, minimum)
		os.Exit(1)
	}

	fmt.Printf("coverage %.2f%% meets required %.2f%%\n", percent, minimum)
}

func parseCoverage(profile string) (covered int, total int) {
	for _, line := range strings.Split(profile, "\n") {
		if line == "" || strings.HasPrefix(line, "mode:") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 3 {
			continue
		}

		statements, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		count, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		total += statements
		if count > 0 {
			covered += statements
		}
	}

	return covered, total
}
