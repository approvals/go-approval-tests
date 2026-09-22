#!/usr/bin/env bash
set -euo pipefail

go build -v ./...
go test ./... -race -coverprofile=coverage.txt -covermode=atomic