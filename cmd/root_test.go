package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Test case 1: Non-existent file
	nonExistentFilePath := "non_existent_file.txt"
	err := readFile(nonExistentFilePath, os.Stdout, false)
	if !os.IsNotExist(err) {
		t.Errorf("Expected error os.ErrNotExist, got: %v", err)
	}

	// Test case 2: Existing file
	existingFilePath := "../test/test1.txt"
	err = readFile(existingFilePath, os.Stdout, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestReadFromStdout(t *testing.T) {
	existingFilePath := "../test/test1.txt"

	var output bytes.Buffer

	// Test case 3: Read from file
	err := readFile(existingFilePath, &output, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	str := "This is a test file #1\n"

	if str != output.String() {
		t.Errorf("Expected %s, got: %s", str, output.String())
	}
}

func TestReadFromStdoutWithNumberedLines(t *testing.T) {
	existingFilePath := "../test/test1.txt"

	var output bytes.Buffer

	// Test case 3: Read from file
	err := readFile(existingFilePath, &output, true)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	str := "1\tThis is a test file #1\n"

	if str != output.String() {
		t.Errorf("Expected %s, got: %s", str, output.String())
	}
}
