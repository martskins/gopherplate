.PHONY: run build

generate:
	go generate ./...

build: generate
	go-bindata out.tmpl
	go build -o gopherplate *.go

install: build
	go install
