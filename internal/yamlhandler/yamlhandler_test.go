package yamlhandler

import (
	"path/filepath"
	"testing"
)

func TestListYAMLFiles(t *testing.T) {

	testDir := filepath.Join("testdata")

	yamlFiles, err := listYAMLFilesInDir(testDir)
	if err != nil {
		t.Fatalf("Error listing YAML files: %v", err)
	}

	expectedFiles := []string{
		filepath.Join(testDir, "file_a.yaml"),
		filepath.Join(testDir, "file_b.yml"),
		filepath.Join(testDir, "file_c.yaml"),
	}

	if len(yamlFiles) != len(expectedFiles) {
		t.Errorf("Incorrect number of YAML files found. Expected %d, got %d", len(expectedFiles), len(yamlFiles))
	}

	for _, expected := range expectedFiles {
		found := false
		for _, file := range yamlFiles {
			if file == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in list of YAML files", expected)
		}
	}
}
