debug:
	 protoc \
	  --plugin=protoc-gen-debug=/home/dev/go/bin/protoc-gen-debug \
	  --proto_path=/home/dev/projects/single/comps/go/api/adapters/grpc/protos \
	  --proto_path=/home/dev/projects/single/comps/go/grpc-plugins/service \
	  --debug_out=".:." \
	  *.proto

generate_validate_pb:
	 protoc \
      --go_out=./pb \
      --proto_path=/home/dev/projects/single/comps/go/api/adapters/grpc/protos \
      --proto_path=/home/dev/projects/single/comps/go/grpc-plugins/service \
      --go_opt=paths=source_relative \
      service.proto

test_plugin_generator:
	 protoc \
      --plugin=protoc-gen-service=/home/dev/projects/single/comps/go/grpc-plugins/service/protoc-gen-service \
      --service_out=./pb \
      --proto_path=/home/dev/projects/single/comps/go/api/adapters/grpc/protos \
      --proto_path=/home/dev/projects/single/comps/go/grpc-plugins/service \
      bank.proto
