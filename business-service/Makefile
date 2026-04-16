GOPATH_BIN := $(shell go env GOPATH)/bin
PROTO_FILES := \
	../proto/business/product.proto \
	../proto/business/customercard.proto

.PHONY: gen gen-proto gen-sqlc test
gen: gen-proto gen-sqlc

gen-proto:
	PATH="$(GOPATH_BIN):$$PATH" protoc \
		-I ../proto \
		--go_out=. \
		--go_opt=module=github.com/handziurdmytro/3JlArODA/business-service \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/handziurdmytro/3JlArODA/business-service \
		$(PROTO_FILES)

gen-sqlc:
	PATH="$(GOPATH_BIN):$$PATH" sqlc generate

test:
	go test ./...
