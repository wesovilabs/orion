#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
# Change into that directory
cd "$DIR"

GOFLAGS=-mod=vendor go run -mod=vendor \
  github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose
