SHELL=/bin/sh

default: run

install:
	go mod download

just-run:
	go run main.go

run: install just-run

test:
	go test ./...

test-force:
	go clean -testcache && go test ./...

test-force-v:
	go clean -testcache && go test -v ./...

test-coverage:
	go test -cover ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy
