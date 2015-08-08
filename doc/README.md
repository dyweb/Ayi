# Ayi Documentation

## Develop

- Setup go workspace following this guide https://golang.org/doc/code.html

````
# This is for ubuntu
# create the workspace folder
mkdir ~/go
mkdir ~/go/bin
mkdir ~/go/pkg

# Add the following lines to ~/.bashrc 
export GOPATH=GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# refresh the file
source ~/.bashrc

````

- clone the project as `$GOPATH/src/github.com/dyweb/Ayi` ( you run git clone command in 
`$GOPATH/src/github.com/dyweb`

- install it as command, in `$GOPATH/src/github.com/dyweb/Ayi` type `go install`
- run `Ayi`, you will some print