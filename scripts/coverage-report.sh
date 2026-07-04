#!/usr/bin/env sh
set -eu

coverage_file="${1:-reports/coverage.out}"
minimum="${2:-80}"
summary_file="reports/coverage-summary.md"
text_file="reports/coverage.txt"
html_file="reports/coverage.html"

mkdir -p reports

go tool cover -func="$coverage_file" > "$text_file"
go tool cover -html="$coverage_file" -o "$html_file"

total_line="$(awk '/^total:/ {print $NF}' "$text_file")"
total_percent="${total_line%\%}"

status="PASS"
awk -v actual="$total_percent" -v min="$minimum" 'BEGIN { exit !(actual < min) }' && status="FAIL"

{
  echo "# Coverage Summary"
  echo
  echo "| Metric | Value |"
  echo "| --- | --- |"
  echo "| Total statement coverage | ${total_line} |"
  echo "| Required minimum | ${minimum}% |"
  echo "| Result | ${status} |"
  echo
  echo "## Function Coverage"
  echo
  echo '```text'
  cat "$text_file"
  echo '```'
  echo
  echo "## Generated Files"
  echo
  echo "- \`coverage.out\`: raw machine-readable Go coverage profile"
  echo "- \`coverage.txt\`: function-level text summary"
  echo "- \`coverage.html\`: visual line-by-line coverage report"
  echo "- \`coverage-summary.md\`: client-readable summary"
} > "$summary_file"

if [ "$status" = "FAIL" ]; then
  echo "coverage ${total_line} is below required ${minimum}%" >&2
  exit 1
fi

echo "coverage reports generated:"
echo "$text_file"
echo "$html_file"
echo "$summary_file"
