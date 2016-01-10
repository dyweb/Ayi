#!/usr/bin/env bash

echo "format use goimports"
goimports -w ./..
echo "lint using golint"
golint ./...