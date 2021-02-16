#!/usr/bin/env bash


for pkg in $(GOFLAGS=-mod=vendor go list -f '{{.Dir}}' ./... | grep -v /vendor/ ); do \
    echo $pkg
    GOFLAGS=-mod=vendor go run -mod=vendor golang.org/x/tools/cmd/goimports -l -w -e $pkg/*.go; \
    GOFLAGS=-mod=vendor go run mvdan.cc/gofumpt -l -w  $pkg/*.go;
done



