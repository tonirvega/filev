package main

import (
	"testing"
)

func TestCreateFileMap(t *testing.T) {

	filesMap, err := createFileMap("tests/fixtures/testdata", []string{})

	if err != nil {

		t.Errorf("Error: %v", err)

	}

	if len(filesMap) != 4 {

		t.Errorf("Expected 4 elements, got %d", len(filesMap))

	}

}
