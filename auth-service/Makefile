.PHONY: gen build run migrate-up migrate-down migrate-status

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

gen:
	protoc --go_out=. --go-grpc_out=. \
		-I ../proto \
		../proto/auth.proto \
		../proto/auth_secret.proto

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server


MIGRATIONS_DIR := ./migrations

# Apply new
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

# Rollback
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

# Show status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status