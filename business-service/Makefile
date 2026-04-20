PWD := $(shell pwd)
PROTO_FILES := \
    ../proto/business/product.proto \
    ../proto/business/category.proto \
    ../proto/business/storeproduct.proto \
    ../proto/business/customercard.proto \
    ../proto/business/employee.proto \
    ../proto/business/check.proto \
    ../proto/business/sale.proto 

.PHONY: gen gen-proto gen-sqlc test
gen: gen-proto gen-sqlc

gen-proto:
	@echo "Generating Proto for business-service via Docker..."
	docker run --rm -v $(PWD)/..:/app -w /app/business-service golang:1.26-alpine sh -c "\
		apk add --no-cache protoc && \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
		export PATH=\"/go/bin:\$$PATH\" && \
		protoc -I ../proto \
		--go_out=. \
		--go_opt=module=github.com/handziurdmytro/3JlArODA/business-service \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/handziurdmytro/3JlArODA/business-service \
		$(PROTO_FILES)"

gen-sqlc:
	@echo "Generating SQLC for business-service via Docker..."
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate

test:
	go test ./...