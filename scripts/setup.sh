#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
# Change into that directory
cd "$DIR"


echo "Downloading dependencies..."
go mod tidy;
go mod vendor;

echo "Create git hooks"
chmod +x .githooks/*
cp .githooks/* .git/hooks/
