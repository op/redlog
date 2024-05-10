#!/usr/bin/make -f

all: toolsinstall tidy build generate

clean:
	@rm -rf ./build

build: build/redlog build/gallery

build/%:
	go build -o $@ ./cmd/$*

generate:
	@build/gallery

tidy:
	@go list -m -f '{{.Dir}}' | xargs -L1 go mod tidy -C

toolsinstall:
	@go install github.com/charmbracelet/freeze@v0.1.6


.PHONY: all build generate tidy toolsinstall
