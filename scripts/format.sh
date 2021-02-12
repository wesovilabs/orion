#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
# Change into that directory
cd "$DIR"

for pkg in $(GOFLAGS=-mod=vendor go list -f '{{.Dir}}' ./... | grep -v /vendor/ ); do \
    GOFLAGS=-mod=vendor go run -mod=vendor golang.org/x/tools/cmd/goimports -local -l -w -e $$pkg/*.go; \
done
