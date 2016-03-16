get-deps:
	glide install
.PHONY: get-deps

travis-install:
	./scripts/travis-install.sh
.PHONY: travis-install

build:
	make get-deps
	echo "build for current platform"
	go build
.PHONY: build

build-osx:
	make get-deps
	echo "build for OS X"
	# FIXME: go build runtime: drawin/amd64 must be bootstrapped using make.bash
	GOOS=drawin GOARCH=386 go build -o Ayi.osx Ayi.go
	mv Ayi.osx build/osx/Ayi
	tar -cvzf build/ayi.osx.tgz build/osx/Ayi
.PHONY: build-osx

build-linux:
	make get-deps
	echo "build for linux"
	GOOS=linux GOARCH=386 CGO_ENABLED=0 godep go build -o Ayi.linux Ayi.go
	mv Ayi.linux build/linux/Ayi
	tar -cvzf build/ayi.linux.tgz build/linux/Ayi
.PHONY: build-linux

build-windows:
	make get-deps
	echo "build for windows"
	GOOS=windows GOARCH=386 godep go build -o Ayi.exe Ayi.go
	mv Ayi.exe build/win/Ayi.exe
	zip build/ayi.win.zip build/win/Ayi.exe
.PHONY: build-windows
	
build-all:
	make get-deps
	make build-osx
	make build-linux
	make build-windows
.PHONY: build-all

test:
	./scripts/test.sh
.PHONY: test

install:
	make get-deps
	go install
.PHONY: install

docker-build-linux:
	docker run --rm -v `pwd`:/go/src/github.com/dyweb/Ayi -e GOPATH=/go:/go/src/github.com/dyweb/Ayi/Godeps/_workspace golang:1.5.3 sh -c "cd /go/src/github.com/dyweb/Ayi && go build -race ."
.PHONY: docker-build-linux
	
