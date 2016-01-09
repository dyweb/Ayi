# Ayi

[![Build Status](https://travis-ci.org/dyweb/Ayi.svg)](https://travis-ci.org/dyweb/Ayi)
[![GoDoc](https://godoc.org/github.com/dyweb/Ayi?status.svg)](https://godoc.org/github.com/dyweb/Ayi)

Centralize all your commands for config development environment

- run several build commands in one command, ie `npm install; composer install;./create_mysql_table.sh;vi /etc/hosts` -> `Ayi install`
- check your environment and try to fix it, ie `Ayi check node` will check if you are using `nvm` and have `node`,`npm`,`gulp` available
- run your test and show result in browser.
- report your machine environment for your colleague to see why your code only runs on your machine. 

## Example

Serve as static server

- `Ayi serve` start a static server in current folder

Config your hosts file

- `Ayi host list` list all your host file
- `Ayi host add -ip 127.0.0.1 -name ayi.dev` add `ayi.dev` to `localhost`
- `Ayi host rm -name ayi.dev` remove `ayi.dev`


