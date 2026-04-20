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
		--go_opt=Mbusiness/category.proto=$(MODULE)/pb/business/category \
		--go_opt=Mbusiness/storeproduct.proto=$(MODULE)/pb/business/storeproduct \
		--go_opt=Mbusiness/customercard.proto=$(MODULE)/pb/business/customercard \
		--go_opt=Mbusiness/employee.proto=$(MODULE)/pb/business/employee \
		--go_opt=Mbusiness/check.proto=$(MODULE)/pb/business/check \
		--go_opt=Mbusiness/sale.proto=$(MODULE)/pb/business/sale \
		--go-grpc_out=. \
		--go-grpc_opt=module=$(MODULE) \
		--go-grpc_opt=Mbusiness/product.proto=$(MODULE)/pb/business/product \
		--go-grpc_opt=Mbusiness/category.proto=$(MODULE)/pb/business/category \
		--go-grpc_opt=Mbusiness/storeproduct.proto=$(MODULE)/pb/business/storeproduct \
		--go-grpc_opt=Mbusiness/customercard.proto=$(MODULE)/pb/business/customercard \
		--go-grpc_opt=Mbusiness/employee.proto=$(MODULE)/pb/business/employee \
		--go-grpc_opt=Mbusiness/check.proto=$(MODULE)/pb/business/check \
		--go-grpc_opt=Mbusiness/sale.proto=$(MODULE)/pb/business/sale \
		../proto/business/product.proto \
		../proto/business/category.proto \
		../proto/business/storeproduct.proto \
		../proto/business/customercard.proto \
		../proto/business/employee.proto \
		../proto/business/check.proto \
		../proto/business/sale.proto

test:
	go test ./...
