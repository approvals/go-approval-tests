#!/bin/bash

go clean -cache
GO_CACHE=off go build -v ./...
GO_CACHE=off go test ./... -race -coverprofile=coverage.txt -covermode=atomic