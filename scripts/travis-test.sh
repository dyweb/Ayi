#!/usr/bin/env bash

# FIXME: should put this in makefile, but I don't know how to make $(glide novendor) works in make file

# switch folder
# get the script path http://stackoverflow.com/questions/4774054/reliable-way-for-a-bash-script-to-get-the-full-path-to-itself
pushd `dirname $0` > /dev/null
SCRIPTPATH=`pwd -P`
popd > /dev/null
# get current working directory
ORIGINAL_WD=${PWD}
# switch to script directory
cd ${SCRIPTPATH}
# switch to parent folder
cd ..

# run the test
# FIXME: the backup seems to have folder issue, will generate a nested fixture folder

echo "backup the fixture"
cp -r fixture fixture-bak    

# TODO: only do this in travis 
# add glide to path
export PATH=$PATH:${SCRIPTPATH}/linux-amd64
# show it is working
glide -v

# enable go vendor feature for 1.5
echo "enable go vendor feature"
export GO15VENDOREXPERIMENT=1

go test -v -cover $(glide novendor)
echo "recover the fixture"
rm -r fixture
mv fixture-bak fixture

# go back to the old working directory
cd ${ORIGINAL_WD}