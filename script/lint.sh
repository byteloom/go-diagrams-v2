#!/usr/bin/env bash
set -e

docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint golangci-lint run --timeout 2m
