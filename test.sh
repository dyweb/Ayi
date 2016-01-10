#!/usr/bin/env bash
#go test -v -cover github.com/dyweb/Ayi/...

#go test -coverprofile=coverage.txt -covermode=atomic

#go test -v -cover github.com/dyweb/Ayi/... -coverprofile=coverage.txt -covermode=atomic

gotestcover -coverprofile=cover.txt -covermode=atomic ./...