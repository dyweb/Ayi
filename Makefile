define AYI_HELP_MSG
Make commands for ayi

help          show help

Dev:
install       install ./cmd/ayi to $$GOPATH/bin/ayi
fmt           gofmt
test          unit test
generate      generate code using gommon

Dev first time:
dep-install   install dependencies based on lock file
dep-update    update dependency based on Gopkg.toml and code

Build:
build         build binary of all platforms and package to zip
endef
export AYI_HELP_MSG

.PHONY: help
help:
	@echo "$$AYI_HELP_MSG"

VERSION = 0.2.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

GO = CGO_ENABLED=0 go
GO_LINUX_BUILD = GOOS=linux GOARCH=amd64
GO_MAC_BUILD = GOOS=darwin GOARCH=amd64
GO_WINDOWS_BUILD = GOOS=windows GOARCH=amd64

.PHONY: fmt
fmt:
	goimports -d -l -w ./cmd ./app ./config ./util

.PHONY: test
test:
	go test -v -cover ./app/... ./config ./util/...

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/ayi

.PHONY: dep-install
dep-install:
	dep ensure -v

.PHONY: dep-update
dep-update:
	dep ensure -update -v

.PHONY: generate
generate:
	gommon generate -v

.PHONY: build build-linux build-mac build-windows
build: build-linux build-mac build-windows

build-linux:
	rm -f ayi-v$(VERSION)-linux-amd64.zip
	$(GO_LINUX_BUILD) go build -ldflags "$(FLAGS)" -o ayi-v$(VERSION)-linux-amd64 ./cmd/ayi
	zip ayi-v$(VERSION)-linux-amd64.zip ayi-v$(VERSION)-linux-amd64
#	rm ayi-v$(VERSION)-linux-amd64

build-mac:
	rm -f ayi-v$(VERSION)-darwin-amd64.zip
	$(GO_MAC_BUILD) go build -ldflags "$(FLAGS)" -o ayi-v$(VERSION)-darwin-amd64 ./cmd/ayi
	zip ayi-v$(VERSION)-darwin-amd64.zip ayi-v$(VERSION)-darwin-amd64
#	rm ayi-v$(VERSION)-darwin-amd64

build-windows:
	rm -f ayi-v$(VERSION)-windows-amd64.zip
	$(GO_WINDOWS_BUILD) go build -ldflags "$(FLAGS)" -o ayi-v$(VERSION)-windows-amd64 ./cmd/ayi
	zip ayi-v$(VERSION)-windows-amd64.zip ayi-v$(VERSION)-windows-amd64
#	rm ayi-v$(VERSION)-windows-amd64

.PHONY: package
package:
	rm -f ayi-v$(VERSION)-linux-amd64.zip
	go build -ldflags "$(FLAGS)" -o ayi-v$(VERSION)-linux-amd64 ./cmd/ayi
	zip ayi-v$(VERSION)-linux-amd64.zip ayi-v$(VERSION)-linux-amd64
	rm ayi-v$(VERSION)-linux-amd64

.PHONY: loc

# https://github.com/Aaronepower/tokei respect ignore file
# exclude playground
loc:
	tokei .