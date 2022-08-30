SHELL=/bin/sh

default: run

install:
	go mod download

just_run:
	go run main.go

run: install just-run

run_http_server:
	go run main.go http_server_run

run_generate_religion:
	go run main.go generate_religion

test:
	go test ./...

test_force:
	go clean -testcache && go test ./...

test_force_v:
	go clean -testcache && go test -v ./...

test_coverage:
	go test -cover ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy
