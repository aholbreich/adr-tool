package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTemplate_Success(t *testing.T) {
	manager := NewTplManager()

	tmpl, err := manager.LoadTemplate()
	if err != nil {
		t.Fatalf("LoadTemplate() failed: %v", err)
	}

	if tmpl == nil {
		t.Error("Expected a parsed template, but got nil")
	}
}

func TestSaveProcessed_Success(t *testing.T) {
	manager := NewTplManager()

	// Sample data for the template
	data := struct {
		Title  string
		Date   string
		Number int
		Status string
	}{
		Title:  "Example ADR",
		Date:   "2025-02-13",
		Number: 1,
		Status: "PROPOSED", // Example status
	}

	// Create a temporary directory for output
	tempDir := t.TempDir()
	outputFile := filepath.Join(tempDir, "output.md")

	// Run the method under test
	err := manager.SaveProcessed(outputFile, data)
	if err != nil {
		t.Fatalf("SaveProcessed() failed: %v", err)
	}

	// Verify the output file exists
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("Output file not created: %v", err)
	}

	// Verify file contents
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if !strings.Contains(string(content), "Example ADR") {
		t.Errorf("Output file does not contain expected title. Got: %s", string(content))
	}

	if !strings.Contains(string(content), "2025-02-13") {
		t.Errorf("Output file does not contain expected date. Got: %s", string(content))
	}
}
