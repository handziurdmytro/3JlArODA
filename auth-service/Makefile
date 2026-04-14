.PHONY: gen build run

gen:
	protoc --go_out=. --go-grpc_out=. \
		-I ../proto \
		../proto/auth.proto \
		../proto/auth_secret.proto

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server