.PHONY: build test clean

BINARY_NAME := goactionctl

build:
	go build -o bin/$(BINARY_NAME)
	mkdir playground

test:
	go test

clean:
	go clean
	rm -rf bin
	rm -rf playground
