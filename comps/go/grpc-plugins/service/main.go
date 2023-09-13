package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

type serviceData struct {
	MessageName     string
	HasCreateMethod bool
}

func main() {
	// Lê os dados de entrada da entrada padrão (stdin)
	logrus.Infof("iniciando plugin")
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Erro ao ler a entrada padrão: %v", err)
	}

	// Parseia os dados de entrada em um CodeGeneratorRequest
	request := &pluginpb.CodeGeneratorRequest{}

	if err := proto.Unmarshal(data, request); err != nil {
		log.Fatalf("Erro ao fazer unmarshal dos dados de entrada: %v", err)
	}

	s := serviceData{}
	for _, proto := range request.GetProtoFile() {
		if proto.GetName() != request.GetFileToGenerate()[0] {
			continue
		}
		for _, service := range proto.GetService() {
			if strings.ToLower(service.GetName()) == "create" {
				s.HasCreateMethod = true
			}
		}

	}
	fmt.Println("files to generate", request.GetFileToGenerate())
	//fmt.Println("service", request.ProtoFile[0].GetEnumType()[0].GetName())

	// Salve os dados de entrada em um arquivo JSON
	outputJSON, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao fazer marshal dos dados de entrada: %v", err)
	}

	// Escreva o JSON em um arquivo de saída
	outputFilePath := "/home/dev/projects/single/comps/go/grpc-plugins/input.json" // Nome do arquivo de saída
	if err := ioutil.WriteFile(outputFilePath, outputJSON, 0644); err != nil {
		log.Fatalf("Erro ao escrever o arquivo de saída: %v", err)
	}

	//for _, p := range request.ProtoFile {
	//	fmt.Println(p.GetName())
	//}

	// Inicializa o protogen com base no request
	//gen := protogen.Options{}

	// Gere o código personalizado com base no request
	//gen.GenerateFiles(request)

	// Gere e escreva o CodeGeneratorResponse
	//response := gen.Response()
	//data, err = proto.Marshal(response)
	//if err != nil {
	//	log.Fatalf("Erro ao fazer marshal da resposta: %v", err)
	//}

	// Escreva a resposta na saída padrão (stdout)
	//if _, err := os.Stdout.Write(data); err != nil {
	//	log.Fatalf("Erro ao escrever a resposta para a saída padrão: %v", err)
	//}
}

// ParseProtoFile reads a .proto file and parses it into a data structure.
func ParseProtoFile(protoFilePath string) (*descriptorpb.FileDescriptorProto, error) {
	// Read the content of the .proto file
	protoData, err := os.ReadFile(protoFilePath)
	if err != nil {
		return nil, err
	}

	// Parse the .proto file
	parser := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(protoData, parser); err != nil {
		return nil, err
	}

	// Return the FileDescriptorProto
	if len(parser.ProtoFile) > 0 {
		return parser.ProtoFile[0], nil
	}

	return nil, fmt.Errorf("no .proto file found in the input file")
}

func GenerateServicePBFile(protoFilePath string) {
	fileDescriptor, err := ParseProtoFile(protoFilePath)
	if err != nil {
		log.Fatal("Error parsing the .proto file: %v", err)
	}
	service := serviceData{}
	// Example of accessing information from the .proto file
	fmt.Printf("File Name: %s\n", fileDescriptor.GetName())
	fmt.Printf("Messages:\n")
	for _, msg := range fileDescriptor.GetMessageType() {
		fmt.Printf("  Message Name: %s\n", msg.GetName())
		service.MessageName = msg.GetName()
		fmt.Printf("  Fields:\n")
		for _, field := range msg.GetField() {
			fmt.Printf("    Field Name: %s\n", field.GetName())
			fmt.Printf("    Field Type: %s\n", field.GetType().String())
		}
	}
}
