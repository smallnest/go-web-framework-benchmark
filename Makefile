.PHONY: build doc fmt lint dev test vet setup bench

# PKG_NAME=$(shell basename `pwd`)
PKG_NAME=gowebbenchmark
CURRENT_DIR=$(shell pwd)
GOPATH := ${CURRENT_DIR}
export GOPATH

setup:
	go get -t -v ./...

build:
	go build -v -o ./$(PKG_NAME)

doc:
	godoc -http=:6060

fmt:
	go fmt ./...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./...

dev:
	DEBUG=* go get && go install && gin -p 8911 -i

test:
	go test ./...

# Runs benchmarks
bench:
	go test ./... -bench=.

# https://godoc.org/golang.org/x/tools/cmd/vet
vet:
	go vet ./...
