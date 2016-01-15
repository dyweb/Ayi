#!/usr/bin/env bash

echo "build for OS X"

# FIXME: go build runtime: drawin/amd64 must be bootstrapped using make.bash
#GOOS=drawin GOARCH=amd64 go build -o Ayi.osx Ayi.go
#GOOS=drawin GOARCH=386 go build -o Ayi.osx Ayi.go
go build -o Ayi.osx Ayi.go

echo "build for windows"
GOOS=windows GOARCH=386 go build -o Ayi.exe Ayi.go
echo "build for linux"
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o Ayi.linux Ayi.go

mv Ayi.osx build/osx/Ayi
mv Ayi.linux build/linux/Ayi
mv Ayi.exe build/win/Ayi.exe

tar -cvzf build/ayi.osx.tgz build/osx/Ayi
tar -cvzf build/ayi.linux.tgz build/linux/Ayi
zip build/ayi.win.zip build/win/Ayi.exe
