#!/usr/bin/env bash
# only need to back fixture when test locally
echo "backup the fixture"
cp -r fixture fixture-bak
go test -v -cover github.com/dyweb/Ayi/...
echo "recover the fixture"
rm -r fixture
mv fixture-bak fixture
