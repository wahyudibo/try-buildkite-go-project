#!/usr/bin/env sh

set -e

golangci-lint run --fix ./... && go test -v -p 1 ./...