.PHONY: fmt lint test build all

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -race ./...

build:
	go build -o bin/app ./cmd/app

all: fmt lint test build