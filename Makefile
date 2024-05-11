#!/usr/bin/make -f

all: toolsinstall tidy check build generate

clean:
	@rm -rf ./build

build: build/redlog build/gallery

build/%:
	go build -o $@ ./cmd/$*

check:
	@go list -m -f '{{.Dir}}' | parallel 'go test {}/...'

coverage:
	go test -coverpkg=./...,./pkg/catppuccin,./internal/logtheme,./internal/themes -covermode=atomic -coverprofile=build/coverage.txt ./... ./pkg/catppuccin/... ./internal/logtheme ./internal/themes

generate:
	@build/gallery

tidy:
	@go list -m -f '{{.Dir}}' | parallel 'go mod tidy -C'

toolsinstall:
	@go install github.com/charmbracelet/freeze@v0.1.6


.PHONY: all build generate tidy toolsinstall
