package system

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(commands ...string) error {
	if len(commands) == 0 {
		return fmt.Errorf("Nenhum comando fornecido")
	}
	var combinedCommand []string
	for _, c := range commands {
		combined := strings.Split(c, " ")
		for idx := range combined {
			combined[idx] = strings.TrimSpace(combined[idx])
			if combined[idx] != "" {
				combinedCommand = append(combinedCommand, combined[idx])
			}
		}
	}

	cmd := exec.Command(combinedCommand[0], combinedCommand[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Erro ao executar o comando: %v", err)
	}

	return nil
}
