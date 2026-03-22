.PHONY: build run test migrate seed tidy

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

migrate:
	go run ./cmd/server -migrate

seed:
	go run ./cmd/server -migrate -seed

tidy:
	go mod tidy
