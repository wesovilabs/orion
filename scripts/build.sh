#!/usr/bin/env bash

# Set module
MODULE="github.com/wesovilabs/orion"
DOCKER_IMG=wesovilabs/orion

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
# Change into that directory
cd "$DIR"

# Get details from commit
COMMIT=$(git log --pretty=format:'%H' -n 1)
VERSION=$(git describe --tags --always)
BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)



# Determine the build mode
BUILD_MODE=${BUILD_MODE:-dev}

# Delete the old dir and crete the new one
echo "==> Removing old directory..."
rm -f bin/*
mkdir -p bin/

# Set LD_FLAGS
LD_FLAGS="-s -w \
 -X ${MODULE}/internal.Commit=${COMMIT} \
 -X ${MODULE}/internal.Version=${VERSION}\
 -X ${MODULE}/internal.BuildDate=${BUILD_DATE}"

# Download dependencies
go mod download

# Build binary
echo "==> Building..."
case ${BUILD_MODE} in
  dev)
    CGO_ENABLED=0 GOFLAGS=-mod=vendor \
      go build \
        -ldflags "${LD_FLAGS}" \
        -o bin/orion ${MODULE}/cmd/orion
    echo
    echo "==> Results:"
    ls -hl bin/
    ;;
  docker)
    GOARCH=amd64 GOOS=linux CGO_ENABLED=0 GOFLAGS=-mod=vendor \
      go build -ldflags "${LD_FLAGS}" -o bin/orion.linux ${MODULE}/cmd/orion
    docker build -f build/docker/Dockerfile -t ${DOCKER_IMG}:local .
    echo
    echo "==> Results:"
    ls -hl bin/
    ;;
  debug)
    CGO_ENABLED=0 GOFLAGS=-mod=vendor \
      go build \
        -gcflags "all=-N -l" \
        -ldflags "${LD_FLAGS}" \
        -o bin/orion ${MODULE}/cmd/orion
      echo
      echo "==> Results:"
      ls -hl bin/
    ;;
  release)
    GOARCH=amd64 GOOS=linux CGO_ENABLED=0 GOFLAGS=-mod=vendor \
      go build -ldflags "${LD_FLAGS}" -o bin/orion.linux ${MODULE}/cmd/orion
	  GOARCH=amd64 GOOS=darwin  CGO_ENABLED=0 GOFLAGS=-mod=vendor \
	    go build -ldflags "${LD_FLAGS}" -o bin/orion.darwin ${MODULE}/cmd/orion
	  GOARCH=amd64 GOOS=windows CGO_ENABLED=0 GOFLAGS=-mod=vendor \
	    go build -ldflags "${LD_FLAGS}" -o bin/orion.exe ${MODULE}/cmd/orion

	  # Remove all tarballs
    rm -r dist
    # Create tarball with all the binaries
    echo "==> Packaging..."
    TMPDIR=$(mktemp -d)
    cp -a bin "${TMPDIR}"; mkdir -p dist
    tar -zcvf dist/orion-"${VERSION}".tar.gz -C "${TMPDIR}" .
    rm -r "${TMPDIR}"
    echo "==> Packaged"
    echo
    echo "==> Results:"
    ls -hl bin/ dist/
    ;;
esac
