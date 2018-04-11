# Ayi

[![GoDoc](https://godoc.org/github.com/dyweb/Ayi?status.svg)](https://godoc.org/github.com/dyweb/Ayi)
[![Go Report Card](https://goreportcard.com/badge/github.com/dyweb/Ayi)](https://goreportcard.com/report/github.com/dyweb/Ayi)
[![Build Status](https://travis-ci.org/dyweb/Ayi.svg)](https://travis-ci.org/dyweb/Ayi)
[![Coverage Status](https://coveralls.io/repos/github/dyweb/Ayi/badge.svg?branch=master)](https://coveralls.io/github/dyweb/Ayi?branch=master)
[![codebeat badge](https://codebeat.co/badges/d45b026a-544e-4f06-8faf-8014e8c02784)](https://codebeat.co/projects/github-com-dyweb-ayi-master)

Centralize all your commands for config development environment

- run several build commands in one command, ie `npm install; composer install` -> `Ayi install`
- run your test and show result in browser. ie `Ayi test`
- wrapper for git, clone to your workspace with multiple url format supported. see [app/git](app/git)
- static server and web ui for a bunch of daily tasks, see [app/web][app/web]
- check your environment and try to fix it, ie `Ayi check node` will check if you are using `nvm` and have `node`,`npm`,`gulp` available
- report your machine environment for your colleague to see why your problems can't be reproduced. 

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
- `go install` this will install Ayi to $GOPATH/bin/Ayi but static assets are not included
- `Ayi install` use the previous compiled Ayi binary to execute install commands in `.ayi.yml`

##### Linux & Mac OS

- `make install`

You will have `Ayi` in your `$GOPATH/bin/Ayi`.
You can also build it using docker. (TODO: cross build using xgo)

`make docker-build-linux` you will get a binary file in the directory. (TODO: the docker build may not be working)

## Usage

Run `Ayi help` to see all available commands.

### Flags

- `-v` for verbose logging
- `-n` for dry run, it will show you the commands, but won't execute them.

### Config 

Ayi will find config file in the following locations and merge them.

- `~/.ayi.yml` a global config for current user, config git hosts, credentials etc.
- `.ayi.yml` in current directory, you should put it in vcs.
- `.ayi.local.yml` in current directly, you should ignore the local config in vcs.

### Autocomplete

NOTE: there is a license problem for including `bash_completion` in `scripts/third_party`

- Use `Ayi bash-gen` to generate autocomplete bash script. (The one is repo might be outdated)
- Add it to your `~/.bashrc` by adding `source path/to/ayi_completion.sh`

- Windows user using `git bash` MUST add `bash_completion` to your `~/.bashr` since it does not ship with git for windows
- ZSH is not supported yet, see https://github.com/dyweb/Ayi/issues/45 for detail

## Development

This project is at early stage and is under heavy development.

- have [glide](https://github.com/Masterminds/glide) installed on you system, we use `vendor` instead of `godeps` or go1.5- style
- run `make get-deps` or `glide install` to pull the dependencies to `vendor` folder, note the `glide.lock` will use fixed version for libraries
- run `go install` to have Ayi installed to your `$GOPATH/bin`
- run `Ayi test` to run tests TODO: take fixture into consideration

- run `glide update` if you have added new packages, NOTE: after import new package in your code, you may need to run `glide update` again, since 
it will analyze the code and ignore subpackages if the package is not used.

## FAQ

- Q: How to change file mode on windows
- A: `git update-index --chmod=+x <file>` from dedek's [answer](http://stackoverflow.com/a/13593391/4116260) for this [question](http://stackoverflow.com/questions/6476513/git-file-permissions-on-windows).

## Contribute

(TODO: put it to contribute.md and use .github for issue and pr template)

1. Before push, use `gofmt` to format your code. (TODO: use git hooks to run fmt before commit (or add?))
2. feel free to add commands and applications, do not make break changes to commands unless necessary.

## About

Ayi is the nick name of a [@dyweb](https://github.com/dyweb) member