#!/usr/bin/env bash

TEST_MODE=${TEST_MODE:-unit}


# Download dependencies
go mod download

case ${TEST_MODE} in
unit)
  GOFLAGS=-mod=vendor go test -p=1  $(GOFLAGS=-mod=vendor go list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{.ImportPath}}{{ end }}' ./... | grep -v test ) -v -timeout 10s;
  ;;
coverage)
  GOFLAGS=-mod=vendor go test -p=1  $(GOFLAGS=-mod=vendor go list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{.ImportPath}}{{ end }}' ./... | grep -v test ) -v -timeout 10s -race -coverprofile=coverage.txt -covermode=atomic;
  ;;
e2e)
  echo "Not implemented yet"
  ;;
esac
