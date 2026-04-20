PWD := $(shell pwd)
MIGRATIONS_DIR := ./migrations

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: gen build run migrate-up migrate-down migrate-status

gen:
	@echo "Generating Proto for auth-service via Docker..."
	docker run --rm -v $(PWD)/..:/app -w /app/auth-service golang:1.26-alpine sh -c "\
		apk add --no-cache protoc && \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
		export PATH=\"/go/bin:\$$PATH\" && \
		protoc --go_out=. --go-grpc_out=. \
		-I ../proto \
		../proto/auth.proto \
		../proto/auth_secret.proto"

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

# Apply new
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

# Rollback
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

# Show status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status