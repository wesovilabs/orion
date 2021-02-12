#!/usr/bin/env bash

for pkg in $(GOFLAGS=-mod=vendor go list -f '{{.Dir}}' ./... | grep -v /vendor/ ); do \
    GOFLAGS=-mod=vendor go run -mod=vendor golang.org/x/tools/cmd/goimports -local -l -w -e $$pkg/*.go; \
done
