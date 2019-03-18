.PHONY: build test run deps

NAME = daydeal

default: test

deps:
	GO111MODULE=on go get ./...

build: deps
	GO111MODULE=on go build -o ${NAME} ./cmd/daydeal

test: deps
	GO111MODULE=on go test -v -race -cover ./...

run: build
	./${NAME}
