#!/usr/bin/env bash

echo "Downloading dependencies..."
go mod tidy;
go mod vendor;

echo "Create git hooks"
