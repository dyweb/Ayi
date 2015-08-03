#!/usr/bin/env bash

# TODO: must paste these code in current terminal, run the sh file won't work
echo "Add GOPATH and update PATH in current shell"
export GOPATH=${PWD}
export PATH=$PATH:$GOPATH/bin
