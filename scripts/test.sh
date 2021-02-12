#!/usr/bin/env bash

TEST_MODE=${TEST_MODE:-unit}

# Remove old report directory
rm -r test_report
mkdir -p test_report

# Download dependencies
go mod download

case ${TEST_MODE} in
unit)
  GOFLAGS=-mod=vendor go test -p=1  $(GOFLAGS=-mod=vendor go list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{.ImportPath}}{{ end }}' ./... | grep -v test ) -v -timeout 10s -short -cover;
  ;;
e2e)
  echo "Not implemented yet"
  ;;
esac
