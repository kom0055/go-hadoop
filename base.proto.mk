PROTO_ROOT_PATH=${GOPATH}/src
#PROTO_ROOT_PATH = $(shell pwd)
#PROTO_ROOT_PATH = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: regenerate
regenerate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
