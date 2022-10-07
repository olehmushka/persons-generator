SHELL=/bin/sh

default: run

install:
	go mod download

run: install run_http_server

run_http_server:
	go run main.go http_server_run

run_generate_religion:
	go run main.go generate_religion

run_generate_world:
	go run main.go generate_world

run_refresh_data:
	go run main.go refresh_data

test:
	go test ./cli/... ./config/... ./core/... ./engine/... ./handlers/... ./internal/...

test_force:
	go clean -testcache && go test ./cli/... ./config/... ./core/... ./engine/... ./handlers/... ./internal/...

test_force_v:
	go clean -testcache && go test -v ./cli/... ./config/... ./core/... ./engine/... ./handlers/... ./internal/...

test_coverage:
	go test -cover ./cli/... ./config/... ./core/... ./engine/... ./handlers/... ./internal/...

fmt:
	go fmt ./cli/... ./config/... ./core/... ./engine/... ./handlers/... ./internal/...

tidy:
	go mod tidy

lint:
	golangci-lint run --timeout 60m
