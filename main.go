package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func processSCSS(inputPath, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	cmd := exec.Command("sass", "--style=compressed", inputPath, outputPath)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	dirInt := "./styles/scss"
	dirOut := "./styles/css"

	err := filepath.Walk(dirInt, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Erro ao acessar %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".scss") {
			relativePath, _ := filepath.Rel(dirInt, path)
			outputPath := filepath.Join(dirOut, strings.Replace(relativePath, ".scss", ".css", 1))
			err := processSCSS(path, outputPath)

			if err != nil {
				fmt.Printf("Erro ao converter %s para CSS: %v\n", path, err)
			} else {
				fmt.Printf("Convertido %s para %s\n", path, outputPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erro ao caminhar pelo diretório:", err)
	} else {
		fmt.Println("Conversão concluída.")
	}
}
