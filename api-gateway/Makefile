GOPATH_BIN := $(shell go env GOPATH)/bin
MODULE := github.com/handziurdmytro/3JlArODA/api-gateway

.PHONY: gen gen-auth gen-business test
gen: gen-auth gen-business

gen-auth:
	PATH="$(GOPATH_BIN):$$PATH" protoc \
		-I ../proto \
		--go_out=. \
		--go-grpc_out=. \
		../proto/auth.proto

gen-business:
	PATH="$(GOPATH_BIN):$$PATH" protoc \
		-I ../proto \
		--go_out=. \
		--go_opt=module=$(MODULE) \
		--go_opt=Mbusiness/product.proto=$(MODULE)/pb/business/product \
		--go_opt=Mbusiness/customercard.proto=$(MODULE)/pb/business/customercard \
		--go-grpc_out=. \
		--go-grpc_opt=module=$(MODULE) \
		--go-grpc_opt=Mbusiness/product.proto=$(MODULE)/pb/business/product \
		--go-grpc_opt=Mbusiness/customercard.proto=$(MODULE)/pb/business/customercard \
		../proto/business/product.proto \
		../proto/business/customercard.proto

test:
	go test ./...
