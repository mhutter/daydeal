.PHONY: build test run deps

NAME = daydeal

default: test

deps:
	go get ./...

build: deps
	go build -o ${NAME} ./cmd/daydeal

test: deps
	go test -v -race -cover ./...

run: build
	./${NAME}
