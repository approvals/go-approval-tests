#!/bin/bash

go build -v ./...
go test ./... -race -coverprofile=coverage.txt -covermode=atomic