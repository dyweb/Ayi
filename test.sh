#!/usr/bin/env bash
go test -v -cover github.com/dyweb/Ayi/...
# TODO: -coverprofile=cover.out

#go test -coverprofile=coverage.txt -covermode=atomic