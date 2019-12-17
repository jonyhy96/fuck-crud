#! /usr/bin/make -f

.PHONY: all
all: package

.PHONY: test
test: 
	go test -tags=unit -timeout 30s -short -v `go list ./...  | grep -v /vendor/`

.PHONY: init
init: 
	pre-commit install

.PHONY: package
package: clean build release

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fuck-crud .

.PHONY: release
release: 
	goreleaser

.PHONY: clean	
clean:
	rm -f dist;
