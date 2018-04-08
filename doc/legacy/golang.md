# Go

## Setup

### Windows

- download the msi installer from https://golang.org/doc/install
- set `GOPATH` 
- add `GOPATH\bin` to `PATH`

You can refer http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/

TODO: may need to set http proxy for go get, ie: https://github.com/cyfdecyf/cow/

## Best practice

- https://golang.org/doc/effective_go.html
- https://talks.golang.org/2013/bestpractices.slide

## Build and Run

- `go run main.go` to run the file directly
- `go build` to build it to current system's binary
- `go install main.go` to install the binary to `$GOPATH\bin` TODO: what if the name duplicate ....
