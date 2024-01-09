package validate

import (
	"bytes"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/spf13/afero"
	"os"
	"testing"
)

func TestModule(t *testing.T) {
	req, err := os.Open("./code_generator_request.pb.bin")
	if err != nil {
		t.Fatal(err)
	}

	fs := afero.NewMemMapFs()
	res := &bytes.Buffer{}

	pgs.Init(
		pgs.ProtocInput(req),  // use the pre-generated request
		pgs.ProtocOutput(res), // capture CodeGeneratorResponse
		pgs.FileSystem(fs),    // capture any custom files written directly to disk
	).RegisterModule(NewModule()).Render()

	// check res and the fs for output
}

//
//func TestGetServicesInfo(t *testing.T) {
//	testCases := []struct {
//		name            string
//		protoFilePath   string
//		expectedServiceCount int
//		expectedOptionValue bool
//		expectedMethodCount int
//	}{
//		{
//			name:            "Teste com arquivo proto válido",
//			protoFilePath:   "banco_test.proto", // Altere conforme necessário
//			expectedServiceCount: 1,
//			expectedOptionValue:  true,
//			expectedMethodCount:  3, // Altere conforme necessário
//		},
//		// Adicione mais casos de teste conforme necessário
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			servicesInfo, err := GetServicesInfo(tc.protoFilePath)
//			if err != nil {
//				t.Errorf("Erro ao obter informações dos serviços: %v", err)
//				return
//			}
//
//			// Verifique o número esperado de serviços
//			if len(servicesInfo) != tc.expectedServiceCount {
//				t.Errorf("Número incorreto de serviços. Esperado: %d, Obtido: %d", tc.expectedServiceCount, len(servicesInfo))
//			}
//
//			// Verifique a opção esperada do serviço (se houver)
//			if tc.expectedOptionValue {
//				if val, ok := servicesInfo[0].Options["kissproject.single.service.crud"]; !ok || val != tc.expectedOptionValue {
//					t.Errorf("Opção 'kissproject.single.service.crud' incorreta. Esperado: %v, Obtido: %v", tc.expectedOptionValue, val)
//				}
//			}
//
//			// Verifique o número esperado de métodos
//			if len(servicesInfo[0].MethodNames) != tc.expectedMethodCount {
//				t.Errorf("Número incorreto de métodos. Esperado: %d, Obtido: %d", tc.expectedMethodCount, len(servicesInfo[0].MethodNames))
//			}
//		})
//	}
//}

// Resto do código (definição de structs e imports) permanece o mesmo
