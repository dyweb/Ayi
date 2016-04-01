# Ayi

[![Go Report Card](https://goreportcard.com/badge/github.com/dyweb/Ayi)](https://goreportcard.com/report/github.com/dyweb/Ayi)
[![Build Status](https://travis-ci.org/dyweb/Ayi.svg)](https://travis-ci.org/dyweb/Ayi)
[![GoDoc](https://godoc.org/github.com/dyweb/Ayi?status.svg)](https://godoc.org/github.com/dyweb/Ayi)

Centralize all your commands for config development environment

- run several build commands in one command, ie `npm install; composer install;./create_mysql_table.sh;vi /etc/hosts` -> `Ayi install`
- check your environment and try to fix it, ie `Ayi check node` will check if you are using `nvm` and have `node`,`npm`,`gulp` available
- run your test and show result in browser.
- report your machine environment for your colleague to see why your code only runs on your machine. 

## Development

This project is at early stage and is not actively developed, better not get into it XD.

- have [glide](https://github.com/Masterminds/glide) installed on you system, we use `vendor` instead of `godeps` or go1.5- style
- run `make get-deps` to pull the dependencies to `vendor` folder, note the `glide.lock` will use fixed version for libraries
- run `go install` to have Ayi installed to your `$GOPATH/bin`
- note, when using `idea golang plugin`, please uncheck the use system GOPATH, since it would took a long time for indexing if you have
a lot of projects in the workspace. related issue https://github.com/dyweb/Ayi/issues/37

**For use idea go plugin**

see https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1820

- `ln -s /home/at15/workspace/src/github.com/dyweb/Ayi/vendor /home/at15/workspace/src/github.com/dyweb/Ayi/vendoor/src`
- add `vendoor` to your project `GOPATH` and the hint should work fine

### Contribute

(TODO: put it to contribute.md and use .github for issue and pr template)

1. Before push, use `gofmt` to format your code. (TODO: use git hooks to run fmt before commit (or add?))
2. feel free to add commands and applications, do not make break changes to commands unless necessary.

## Quick Start

### Installation

#### Use binary

- download the binary from the [release page](https://github.com/dyweb/Ayi/releases)
- run it directly or put it in you path, like `mv Ayi /usr/loca/bin/Ayi`

#### From source

- clone the project `git clone git@github.com:dyweb/Ayi.git`
- move to the directory `cd Ayi`
- install `glide`, see https://github.com/Masterminds/glide (TODO: put download glide to makefile)
- run `make install`, you will have `Ayi` in your `$GOPATH/bin/Ayi`

You can also build it using docker (TODO: cross build using xgo)

`make docker-build-linux` you will get a binary file in the directory. (FIXME: it is still using godeps)

### Usage

```
NAME:
   Ayi - Let Ayi do it for you

USAGE:
   Ayi [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   mie, arrowrowe	roast mie
   util-dummy		dummy util command
   static		serve static files
   hosts, host		config/show  host
   git, g		git command wrapper
   mail, m		send mail to all web stuff
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

#### Example

Serve as static server

- `Ayi static` start a static server in current folder

Config your hosts file

- `Ayi host list` list all your host file
- `Ayi host add -ip 127.0.0.1 -name ayi.dev` add `ayi.dev` to `localhost`
- `Ayi host rm -name ayi.dev` remove `ayi.dev`