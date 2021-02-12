#!/usr/bin/env bash

GOFLAGS=-mod=vendor go run -mod=vendor \
  github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose
