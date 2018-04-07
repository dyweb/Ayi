VERSION = 0.2.0
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/ayi

.PHONY: update-dep
update-dep:
	dep ensure -update

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./ayi

.PHONY: test
test:
	go test -v -cover ./ayi/...