PROTO_DIR := protos
OUT_DIR := pb

# Comando para gerar os arquivos .cs a partir dos arquivos .proto
generate_proto: generate_proto_client
	@mkdir -p $(OUT_DIR)
	#  @protoc --csharp_out=$(OUT_DIR) --grpc_out=$(OUT_DIR) --plugin=protoc-gen-grpc=/usr/bin/grpc_csharp_plugin $(PROTO_DIR)/*.proto
	
generate_proto_client:
	@protoc --go_out=../clients/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=../clients/go $(PROTO_DIR)/*.proto

# Comando para instalar o grpc_csharp_plugin
install-plugin:
	sudo apt install -y protobuf-compiler-grpc

# Comando para limpar os arquivos gerados
clean:
	@rm -rf $(OUT_DIR)
