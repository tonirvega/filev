package yamlhandler

import (
	"os"
	"path/filepath"
)

func listYAMLFilesInDir(rootDir string) ([]string, error) {
	var yamlFiles []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			yamlFiles = append(yamlFiles, path)
		}
		return nil
	})

	return yamlFiles, err
}
