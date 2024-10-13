package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func createFileMap(rootDir string) (map[string]string, error) {

	filesMap := make(map[string]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			return err
		}

		if !info.IsDir() {

			body, err := os.ReadFile(path)

			if err != nil {

				return err

			}

			filesMap[path] = string(body)
		}

		return nil

	})

	if err != nil {

		return nil, err

	}

	return filesMap, nil

}

func openWithVim(filePath string) {
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
