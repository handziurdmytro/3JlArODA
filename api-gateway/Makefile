.PHONY: gen
gen:
	mkdir -p pb
	protoc --go_out=. --go-grpc_out=. \
	-I ../proto \
	../proto/auth.proto \