.PHONY: run build

run: build
	./sqlgenx \
		-src ~/go/src/github.com/martskins/restoo/main.go \
		-target Item

build:
	go-bindata out.tmpl
	go build .

install: build
	go install
