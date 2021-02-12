#!/usr/bin/env bash
# Set module
MODULE="github.com/wesovilabs/orion"

RUN_MODE=${RUN_MODE:-code}

case ${RUN_MODE} in
code)
  GOFLAGS=-mod=vendor go  run ${MODULE}/cmd/orion;
  ;;
debug)
  echo "running code in debug mode"
  GOFLAGS=-mod=vendor go run github.com/go-delve/delve/cmd/dlv debug \
    --headless --api-version=2 --log --listen=127.0.0.1:2345 ${MODULE}/cmd/orion
  ;;
esac
