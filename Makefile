.PHONY: build run test lint

build:
	go build -o ipocalc ./cmd/server

run:
	./ipocalc

test:
	go test ./... -cover

lint:
	golangci-lint run
