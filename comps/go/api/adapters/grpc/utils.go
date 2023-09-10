package grpc

import (
	"fmt"
	"github.com/kissprojects/single/comps/go/system"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isPluginInstalled(pluginPath string) bool {
	cmd := exec.Command("go", "list", "-f", pluginPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.ReplaceAll(string(output), "\n", "") == pluginPath
}

func InstallProtocPlugins() {
	plugins := []string{
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
	}
	for _, plugin := range plugins {
		if isPluginInstalled(plugin) {
			log.Debugf("Plugin %s is already installed.\n", plugin)
			continue
		}

		cmd := exec.Command("go", "install", plugin)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("erro while install the plugin %s: %v\n", plugin, err)
		} else {
			fmt.Printf("Plugin %s installed with success.\n", plugin)
		}
	}
}

// generatePBFiles generate pb files based on all .proto files
func generatePBFiles() error {

	executablePath, err := os.Executable()
	if err != nil {
		log.Errorf("error getting the path of the executable: %v", err)
	}

	rootDir := filepath.Dir(executablePath)

	// Anonymous function to process the .proto files in a directory
	var processDir func(string) error
	processDir = func(dir string) error {
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				// Recursively processes subdirectories
				if err := processDir(filepath.Join(dir, file.Name())); err != nil {
					return err
				}
			} else if strings.HasSuffix(file.Name(), ".proto") {
				// Found a .proto file
				protoPath := filepath.Join(dir, file.Name())

				// Output directory for .pb.go files
				pbDir := filepath.Join(dir, "pb")

				// Create the "pb" directory if it doesn't exist
				if _, err := os.Stat(pbDir); os.IsNotExist(err) {
					if err := os.Mkdir(pbDir, 0755); err != nil {
						return err
					}
				}

				// Command to generate .pb.go files with gRPC and regular protobuf support
				err = system.RunCommand("protoc",
					"--grpc-gateway_out=:"+pbDir,               // Use --go-grpc_out for gRPC
					"--grpc-gateway_opt paths=source_relative", // Option to keep paths relative for gRPC
					"--go_out=:"+pbDir,                         // Use --go_out for regular protobuf
					"--go_opt=paths=source_relative",           // Option to keep paths relative to regular protobuf
					"--proto_path=/home/dev/projects/single/comps/go/api/adapters/grpc/protos",
					"--proto_path="+dir,
					protoPath,
				)
				if err != nil {
					return err
				}

				// generate rest routers using grpc-gateway
				err = system.RunCommand("protoc -I . ",
					"--grpc-gateway_out=:"+pbDir,
					"--grpc-gateway_opt paths=source_relative",
					"--proto_path=/home/dev/projects/single/comps/go/api/adapters/grpc/protos",
					"--openapiv2_out "+pbDir,
					"--openapiv2_opt logtostderr=true",
					"--proto_path="+dir,
					protoPath,
				)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	return processDir(rootDir)
}

func generateSwagger() error {
	executablePath, err := os.Executable()
	if err != nil {
		log.Errorf("error getting the path of the executable: %v", err)
	}

	rootDir := filepath.Dir(executablePath)
	files, err := FindSwaggerJSONFiles(rootDir)
	if err != nil {

	}

	return GenerateSwaggerHTML(rootDir+"/swagger", files...)
}

// FindSwaggerJSONFiles finds all files in the base directory that have ".swagger.json" in their names.
func FindSwaggerJSONFiles(baseDir string) ([]string, error) {
	var swaggerFiles []string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.Contains(info.Name(), ".swagger.json") {
			swaggerFiles = append(swaggerFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return swaggerFiles, nil
}

// GenerateSwaggerHTML generates an HTML Swagger file based on multiple JSON files.
func GenerateSwaggerHTML(outputPath string, jsonFiles ...string) error {
	// Create the "pb" directory if it doesn't exist
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.Mkdir(outputPath, 0755); err != nil {
			return err
		}
	}

	// Create a command to generate the Swagger HTML
	cmd := exec.Command("swagger", "flatten", "-o", outputPath)
	cmd.Args = append(cmd.Args, jsonFiles...)

	// Configure the output to an HTML file
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error generating Swagger HTML: %v", err)
	}

	return nil
}
