include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

deps:
	@ echo
	@ echo "Starting downloading dependencies..."
	@ echo
	@ go get -u ./...

server:
	@ echo
	@ echo "Starting the server..."
	@ echo
	@ MONITOR_PORT=${SERVER_MONITOR_PORT} PORT=$(SERVER_PORT) go run ./cmd/server/main.go

test:
	@ echo
	@ echo "Starting running tests..."
	@ echo
	@ go test -cover ./...

debug:
	@ echo
	@ echo "Starting the debug..."
	@ echo
	@ MONITOR_PORT=3100 go run ./cmd/debug

%:
	@: