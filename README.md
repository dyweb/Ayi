# Ayi

[![Go Report Card](https://goreportcard.com/badge/github.com/dyweb/Ayi)](https://goreportcard.com/report/github.com/dyweb/Ayi)
[![Build Status](https://travis-ci.org/dyweb/Ayi.svg)](https://travis-ci.org/dyweb/Ayi)
[![GoDoc](https://godoc.org/github.com/dyweb/Ayi?status.svg)](https://godoc.org/github.com/dyweb/Ayi)

Centralize all your commands for config development environment

- run several build commands in one command, ie `npm install; composer install;./create_mysql_table.sh;vi /etc/hosts` -> `Ayi install`
- check your environment and try to fix it, ie `Ayi check node` will check if you are using `nvm` and have `node`,`npm`,`gulp` available
- run your test and show result in browser.
- report your machine environment for your colleague to see why your code only runs on your machine. 

## Quick Start

### Installation

#### The Go way

Just run the follow code in the terminal:

```
go build github.com/dyweb/Ayi
```

#### The Docker way

If you don't want to install golang in your PC or Mac, just run:

```
make docker-build-linux
```

you will get a binary file in the directory.

#### The Github way

you could get the binary in the [release page](https://github.com/dyweb/Ayi/releases)

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


## For Dev

1. Before push, use `gofmt` to format your code.
2. feel free to add commands, do not change commands.
