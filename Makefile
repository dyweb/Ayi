get-deps:
	glide install
.PHONY: get-deps

godep-build:
	echo "build for OS X"

	# FIXME: go build runtime: drawin/amd64 must be bootstrapped using make.bash
	#GOOS=drawin GOARCH=amd64 go build -o Ayi.osx Ayi.go
	#GOOS=drawin GOARCH=386 go build -o Ayi.osx Ayi.go
	godep go build -o Ayi.osx Ayi.go

	echo "build for windows"
	GOOS=windows GOARCH=386 godep go build -o Ayi.exe Ayi.go
	echo "build for linux"
	GOOS=linux GOARCH=386 CGO_ENABLED=0 godep go build -o Ayi.linux Ayi.go

	mv Ayi.osx build/osx/Ayi
	mv Ayi.linux build/linux/Ayi
	mv Ayi.exe build/win/Ayi.exe

	tar -cvzf build/ayi.osx.tgz build/osx/Ayi
	tar -cvzf build/ayi.linux.tgz build/linux/Ayi
	zip build/ayi.win.zip build/win/Ayi.exe
.PHONY: godep-build

godep-build-osx:
	echo "build for OS X"

	# FIXME: go build runtime: drawin/amd64 must be bootstrapped using make.bash
	#GOOS=drawin GOARCH=amd64 go build -o Ayi.osx Ayi.go
	#GOOS=drawin GOARCH=386 go build -o Ayi.osx Ayi.go
	godep go build -o Ayi.osx Ayi.go
	mv Ayi.osx build/osx/Ayi
	tar -cvzf build/ayi.osx.tgz build/osx/Ayi
.PHONY: godep-build-osx

godep-build-linux:
	echo "build for linux"
	GOOS=linux GOARCH=386 CGO_ENABLED=0 godep go build -o Ayi.linux Ayi.go
	mv Ayi.linux build/linux/Ayi
	tar -cvzf build/ayi.linux.tgz build/linux/Ayi
.PHONY: godep-build-linux

godep-build-windows:
	echo "build for windows"
	GOOS=windows GOARCH=386 godep go build -o Ayi.exe Ayi.go
	mv Ayi.exe build/win/Ayi.exe
	zip build/ayi.win.zip build/win/Ayi.exe
.PHONY: godep-build-windows
	
naive-build:
	make get-deps
	echo "build for OS X"

	# FIXME: go build runtime: drawin/amd64 must be bootstrapped using make.bash
	#GOOS=drawin GOARCH=amd64 go build -o Ayi.osx Ayi.go
	#GOOS=drawin GOARCH=386 go build -o Ayi.osx Ayi.go
	go build -o Ayi.osx Ayi.go

	echo "build for windows"
	GOOS=windows GOARCH=386 go build -o Ayi.exe Ayi.go
	echo "build for linux"
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o Ayi.linux Ayi.go

	mv Ayi.osx build/osx/Ayi
	mv Ayi.linux build/linux/Ayi
	mv Ayi.exe build/win/Ayi.exe

	tar -cvzf build/ayi.osx.tgz build/osx/Ayi
	tar -cvzf build/ayi.linux.tgz build/linux/Ayi
	zip build/ayi.win.zip build/win/Ayi.exe
.PHONY: naive-build

# TODO: migrate to use glide and merge with test-local
test-travis:
	go test -v ./...
.PHONY: test-travis

test-local:
	./scripts/test.sh
.PHONY: test-local

docker-build-linux:
	docker run --rm -v `pwd`:/go/src/github.com/dyweb/Ayi -e GOPATH=/go:/go/src/github.com/dyweb/Ayi/Godeps/_workspace golang:1.5.3 sh -c "cd /go/src/github.com/dyweb/Ayi && go build -race ."
.PHONY: docker-build-linux
	
