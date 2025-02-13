package config

import (
	"adr-tool/internal/model"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestLoadConfig checks if configuration is loaded correctly from a file
func TestLoadConfig(t *testing.T) {
	tempDir := t.TempDir()
	testConfigPath := filepath.Join(tempDir, "config.json")

	// Create a temporary config file
	testConfig := model.AdrConfig{
		BaseDir:    "test/base/dir",
		CurrentAdr: 42,
	}
	bytes, _ := json.MarshalIndent(testConfig, "", " ")
	os.WriteFile(testConfigPath, bytes, 0644)

	// Override the default config path for testing
	resolver := PathResolverInst()
	resolver.BaseDir = tempDir

	// Load the config using the manager
	manager := NewConfigManager()
	config, err := manager.LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	// Check if the loaded values match
	if config.BaseDir != testConfig.BaseDir {
		t.Errorf("BaseDir = %s; want %s", config.BaseDir, testConfig.BaseDir)
	}
	if config.CurrentAdr != testConfig.CurrentAdr {
		t.Errorf("CurrentAdr = %d; want %d", config.CurrentAdr, testConfig.CurrentAdr)
	}
}

// TestUpdateConfig checks if updating and saving config works correctly
func TestUpdateConfig(t *testing.T) {
	tempDir := t.TempDir()
	testConfigPath := filepath.Join(tempDir, "config.json")

	// Override the default config path for testing
	resolver := PathResolverInst()
	resolver.BaseDir = tempDir

	// Initialize the manager and a new config object
	manager := NewConfigManager()
	newConfig := model.AdrConfig{
		BaseDir:    "new/base/dir",
		CurrentAdr: 99,
	}

	// Update the config
	err := manager.UpdateConfig(newConfig)
	if err != nil {
		t.Fatalf("UpdateConfig() failed: %v", err)
	}

	// Read the file and check its content
	bytes, err := os.ReadFile(testConfigPath)
	if err != nil {
		t.Fatalf("Failed to read updated config file: %v", err)
	}

	var savedConfig model.AdrConfig
	json.Unmarshal(bytes, &savedConfig)

	if savedConfig.BaseDir != newConfig.BaseDir {
		t.Errorf("BaseDir = %s; want %s", savedConfig.BaseDir, newConfig.BaseDir)
	}
	if savedConfig.CurrentAdr != newConfig.CurrentAdr {
		t.Errorf("CurrentAdr = %d; want %d", savedConfig.CurrentAdr, newConfig.CurrentAdr)
	}
}
