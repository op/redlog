#!/usr/bin/make -f

all: toolsinstall format test build generate

clean: ## clean up
	@rm -rf ./build

go.work: ## initiate go.work
	go work init
	go work use -r .

build: go.work build/redlog build/gallery ## build

build/%:
	go build -o $@ ./internal/cmd/$*

test: ## run tests
	go list -m -f '{{.Dir}}' | parallel 'go test {}/...'

coverage: ## generate coverage report
	@mkdir -p build/
	go test -coverpkg=./...,./pkg/catppuccin,./internal/logtheme,./internal/themes -covermode=atomic -coverprofile=build/coverage.txt ./... ./pkg/catppuccin/... ./internal/logtheme ./internal/themes

generate: build/gallery ## generate images in asset/
	@build/gallery

format: ## format source code
	@go list -m -f '{{.Dir}}' | parallel -v 'go mod tidy -C'

toolsinstall: ## install tools
	@go install github.com/charmbracelet/freeze@v0.1.6

help:
	@grep -Eh '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[35m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all clean build coverage format generate help test toolsinstall
