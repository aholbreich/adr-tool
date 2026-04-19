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

func TestTemplateFilePath(t *testing.T) {
	resolver := PathResolverInst()
	expected := filepath.Join(resolver.ConfigFolderPath(), DefaultTemplateFileName)
	if resolver.TemplateFilePath() != expected {
		t.Errorf("TemplateFilePath() = %s; want %s", resolver.TemplateFilePath(), expected)
	}
}

// TestIsFilepathGitRepo checks .git detection
func TestIsFilepathGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("create .git dir: %v", err)
	}

	resolver := &PathResolver{BaseDir: tempDir}

	if !resolver.IsFilepathGitRepo() {
		t.Error("IsFilepathGitRepo() = false; expected true")
	}

	if err := os.RemoveAll(gitDir); err != nil {
		t.Fatalf("remove .git dir: %v", err)
	}
	if resolver.IsFilepathGitRepo() {
		t.Error("IsFilepathGitRepo() = true; expected false")
	}
}

func TestIsFilepathGitRepoUsesBaseDirInsteadOfProcessCwd(t *testing.T) {
	tempDir := t.TempDir()
	if err := os.Mkdir(filepath.Join(tempDir, ".git"), 0755); err != nil {
		t.Fatalf("create .git dir: %v", err)
	}

	otherDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer os.Chdir(originalDir)
	if err := os.Chdir(otherDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	resolver := &PathResolver{BaseDir: tempDir}
	if !resolver.IsFilepathGitRepo() {
		t.Fatal("expected IsFilepathGitRepo() to use BaseDir rather than current process directory")
	}
}
