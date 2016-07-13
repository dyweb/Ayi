# Ayi

[![GoDoc](https://godoc.org/github.com/dyweb/Ayi?status.svg)](https://godoc.org/github.com/dyweb/Ayi)
[![Go Report Card](https://goreportcard.com/badge/github.com/dyweb/Ayi)](https://goreportcard.com/report/github.com/dyweb/Ayi)
[![Build Status](https://travis-ci.org/dyweb/Ayi.svg)](https://travis-ci.org/dyweb/Ayi)
[![Coverage Status](https://coveralls.io/repos/github/dyweb/Ayi/badge.svg?branch=master)](https://coveralls.io/github/dyweb/Ayi?branch=master)

Centralize all your commands for config development environment

- run several build commands in one command, ie `npm install; composer install;./create_mysql_table.sh;vi /etc/hosts` -> `Ayi install`
- check your environment and try to fix it, ie `Ayi check node` will check if you are using `nvm` and have `node`,`npm`,`gulp` available
- run your test and show result in browser. ie `Ayi test`
- report your machine environment for your colleague to see why your problems can't be reproduced. 

## Development

This project is at early stage and is under heavy development.

- have [glide](https://github.com/Masterminds/glide) installed on you system, we use `vendor` instead of `godeps` or go1.5- style
- run `make get-deps` or `glide install` to pull the dependencies to `vendor` folder, note the `glide.lock` will use fixed version for libraries
- run `go install` to have Ayi installed to your `$GOPATH/bin`
- run `go test -v -cover $(glide novendor)` to run tests TODO: take fixture into consideration

- run `glide update` if you have added new packages, NOTE: after import new package in your code, you may need to run `glide update` again, since 
it will analyze the code and ignore subpackages if the package is not used.

### FAQ

- Q: How to change file mode on windows
- A: `git update-index --chmod=+x <file>` from dedek's [answer](http://stackoverflow.com/a/13593391/4116260) for this [question](http://stackoverflow.com/questions/6476513/git-file-permissions-on-windows).

### Contribute

(TODO: put it to contribute.md and use .github for issue and pr template)

1. Before push, use `gofmt` to format your code. (TODO: use git hooks to run fmt before commit (or add?))
2. feel free to add commands and applications, do not make break changes to commands unless necessary.

## Quick Start

### Installation

#### Use binary

There are NO stable release now, better build from source

- download the binary from the [release page](https://github.com/dyweb/Ayi/releases)
- run it directly or put it in you path, like `mv Ayi /usr/loca/bin/Ayi`

#### From source

- clone the project `git clone git@github.com:dyweb/Ayi.git`
- move to the directory `cd Ayi`
- install `glide`, see https://github.com/Masterminds/glide (TODO: put download glide to makefile)

##### Windows

- `glide install`
- `go install`

##### Linux & Mac OS

- `make install`

You will have `Ayi` in your `$GOPATH/bin/Ayi`.
You can also build it using docker. (TODO: cross build using xgo)

`make docker-build-linux` you will get a binary file in the directory. (TODO: the docker build may not be working)

### Usage

```
Ayi is a collection of small applications and tools that speed up your develop process

Usage:
  Ayi [flags]
  Ayi [command]

Available Commands:
  test        run test configed in .ayi.yml
  version     current version of Ayi

Flags:
      --config string   config file (default is $HOME/.ayi.yaml)
  -t, --toggle          Help message for toggle
  -v, --verbose         verbose output
      --version         show current version

Use "Ayi [command] --help" for more information about a command.
```

In order to use Ayi for a project, you should copy `.ayi.example.yml` and rename it to `.ayi.yml`, put it in project root.