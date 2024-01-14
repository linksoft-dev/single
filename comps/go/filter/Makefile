.PHONY: generate

generate:
	protoc --go_out==source_relative:. --go-grpc_out==source_relative:. --proto_path=. --proto_path=../grpc/protos filter.proto