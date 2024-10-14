package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func createFileMap(rootDir string, ignoredPaths []string) (map[string]string, error) {

	filesMap := make(map[string]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			return err
		}

		if !info.IsDir() && !isIgnoredPath(path, ignoredPaths) && isNotASymlink(path) {

			body, err := os.ReadFile(path)

			if err != nil {

				return err

			}
			filesSize[path] = info.Size()
			filesMap[path] = string(body)
		}

		return nil

	})

	if err != nil {

		return nil, err

	}

	return filesMap, nil

}

func isIgnoredPath(path string, ignoredPaths []string) bool {
	for _, ignoredPath := range ignoredPaths {
		if strings.Contains(path, ignoredPath) {
			return true
		}
	}
	return false
}

func matchesFilter(path, filter, content string) bool {

	filter = sanitizeString(filter)

	content = sanitizeString(content)

	return strings.Contains(
		strings.ToLower(path),
		strings.ToLower(filter),
	) || strings.Contains(
		strings.ToLower(content),
		strings.ToLower(filter),
	) || filter == ""
}

func sanitizeString(s string) string {
	return strings.TrimSpace(s)
}

func openWithVim(filePath string) {
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func isNotASymlink(path string) bool {
	info, err := os.Lstat(path)

	if err != nil {

		return false

	}

	return info.Mode()&os.ModeSymlink == 0

}
