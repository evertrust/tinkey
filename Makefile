.PHONY: build test pre-run

SHELL := /bin/bash

UNAME_S := $(shell uname -s)

REPO_NAME := tinkey



test:
	go fmt $$(go list ./... | grep -v /vendor/)
	go vet $$(go list ./... | grep -v /vendor/)
	go test -race $(go list ./... | grep -v /vendor/)

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'" -trimpath -o ./build/tinkey-linux-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "-extldflags '-static'" -trimpath -o ./build/tinkey-linux-arm64
	GOOS=windows GOARCH=amd64 go build -ldflags "-extldflags '-static'" -trimpath -o ./build/tinkey.exe
	GOOS=darwin GOARCH=arm64 go build -ldflags "-extldflags '-static'" -trimpath -o ./build/tinkey-darwin-arm64
	GOOS=darwin GOARCH=amd64 go build -ldflags "-extldflags '-static'" -trimpath -o ./build/tinkey-darwin-amd64