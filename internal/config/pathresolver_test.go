package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSingletonBehavior(t *testing.T) {
	resolver1 := PathResolverInst()
	resolver2 := PathResolverInst()

	if resolver1 != resolver2 {
		t.Error("PathResolverInst() did not return the same instance")
	}
}

func TestConfigFolderPath(t *testing.T) {
	resolver := PathResolverInst()
	expected := filepath.Join(resolver.BaseDir, DefaultConfigFolderName)
	if resolver.ConfigFolderPath() != expected {
		t.Errorf("ConfigFolderPath() = %s; want %s", resolver.ConfigFolderPath(), expected)
	}
}

func TestConfigFilePath(t *testing.T) {
	resolver := PathResolverInst()
	expected := filepath.Join(resolver.ConfigFolderPath(), DefaultConfigFileName)
	if resolver.ConfigFilePath() != expected {
		t.Errorf("ConfigFilePath() = %s; want %s", resolver.ConfigFilePath(), expected)
	}
}

func TestTemplateFilePath(t *testing.T) {
	resolver := PathResolverInst()
	expected := filepath.Join(resolver.ConfigFolderPath(), DefaultTemplateFileName)
	if resolver.TemplateFilePath() != expected {
		t.Errorf("TemplateFilePath() = %s; want %s", resolver.TemplateFilePath(), expected)
	}
}

// TestIsFilepathGitRepo checks .git detection
func TestIsFilepathGitRepo(t *testing.T) {

	resolver := PathResolverInst()

	// Create a temporary directory to simulate a git repo
	tempDir := t.TempDir()
	gitDir := filepath.Join(tempDir, ".git")
	os.Mkdir(gitDir, 0755)

	// Change to the temporary directory
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tempDir)

	if !resolver.IsFilepathGitRepo() {
		t.Error("IsFilepathGitRepo() = false; expected true")
	}

	// Remove the .git directory and retest
	os.RemoveAll(gitDir)
	if resolver.IsFilepathGitRepo() {
		t.Error("IsFilepathGitRepo() = true; expected false")
	}
}
