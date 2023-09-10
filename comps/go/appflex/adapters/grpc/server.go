package grpc

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func New(port string) Adapter {
	return Adapter{port: port}
}

type appInterface interface {
	Register(grpcServer *grpc.Server)
}

// Adapter struct to save apps that implement grpc adapters and set port that grpc server should run
type Adapter struct {
	port string
	apps []appInterface
}

func (g *Adapter) Add(app appInterface) {
	g.apps = append(g.apps, app)
}

// Run method that implements Adapter interface
func (g Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", g.port, err)
	}
	grpcServer := grpc.NewServer()
	for _, app := range g.apps {
		app.Register(grpcServer)
	}
	reflection.Register(grpcServer)
	log.Infof("GRPC server listening on %s\n", g.port)
	go generatePBFiles()
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC over %s: %v", g.port, err)
	}
}

func generatePBFiles() error {

	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Erro ao obter o caminho do executável:", err)
	}

	// Obtém o diretório pai do caminho do executável
	rootDir := filepath.Dir(executablePath)

	// Função anônima para processar os arquivos .proto em um diretório
	var processDir func(string) error
	processDir = func(dir string) error {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				// Recursivamente processa subdiretórios
				if err := processDir(filepath.Join(dir, file.Name())); err != nil {
					return err
				}
			} else if strings.HasSuffix(file.Name(), ".proto") {
				// Encontrou um arquivo .proto
				protoPath := filepath.Join(dir, file.Name())

				// Diretório de saída para os arquivos .pb.go
				pbDir := filepath.Join(dir, "pb")

				// Cria o diretório "pb" se não existir
				if _, err := os.Stat(pbDir); os.IsNotExist(err) {
					if err := os.Mkdir(pbDir, 0755); err != nil {
						return err
					}
				}

				// Comando para gerar os arquivos .pb.go com suporte a gRPC e protobuf regular
				cmd := exec.Command(
					"protoc",
					"--go-grpc_out=:"+pbDir,               // Use --go-grpc_out para gRPC
					"--go-grpc_opt=paths=source_relative", // Opção para manter os caminhos relativos para gRPC
					"--go_out=:"+pbDir,                    // Use --go_out para protobuf regular
					"--go_opt=paths=source_relative",      // Opção para manter os caminhos relativos para protobuf regular
					"--proto_path="+dir,                   // Use o diretório atual como proto_path
					protoPath,
				)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				fmt.Printf("Gerando arquivos .pb.go para gRPC e protobuf regular: %s\n", protoPath)

				if err := cmd.Run(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return processDir(rootDir)
}
