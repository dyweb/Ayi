get-deps:
	go get ./...
.PHONY: get-deps

tests:
	go test -v ./...
.PHONY: tests