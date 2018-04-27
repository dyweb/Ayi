VERSION = 0.2.0
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./app ./config ./util

.PHONY: test
test:
	go test -v -cover ./app/... ./config ./util/...

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/ayi

.PHONY: update-dep
update-dep:
	dep ensure -update

.PHONY: generate
generate:
	gommon generate -v

.PHONY: package
package:
	rm -f ayi-v$(VERSION)-linux-amd64.zip
	go build -ldflags "$(FLAGS)" -o ayi-v$(VERSION)-linux-amd64 ./cmd/ayi
	zip ayi-v$(VERSION)-linux-amd64.zip ayi-v$(VERSION)-linux-amd64
	rm ayi-v$(VERSION)-linux-amd64
