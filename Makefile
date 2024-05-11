#!/usr/bin/make -f

all: toolsinstall tidy test build generate

clean:
	@rm -rf ./build

go.work:
	go work init
	go work use -r .

build: go.work build/redlog build/gallery

build/%:
	go build -o $@ ./internal/cmd/$*

test:
	go list -m -f '{{.Dir}}' | parallel 'go test {}/...'

coverage:
	@mkdir -p build/
	go test -coverpkg=./...,./pkg/catppuccin,./internal/logtheme,./internal/themes -covermode=atomic -coverprofile=build/coverage.txt ./... ./pkg/catppuccin/... ./internal/logtheme ./internal/themes

generate:
	@build/gallery

tidy:
	@go list -m -f '{{.Dir}}' | parallel -v 'go mod tidy -C'

toolsinstall:
	@go install github.com/charmbracelet/freeze@v0.1.6


.PHONY: all clean build coverage generate tidy test toolsinstall
