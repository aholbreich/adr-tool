package config

import (
	"adr-tool/internal/model"
	"encoding/json"
	"fmt"
	"os"
)

type ConfigManager struct{}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

func (cm *ConfigManager) LoadConfig() (model.AdrConfig, error) {
	var config model.AdrConfig
	bytes, err := os.ReadFile(PathResolverInst().ConfigFilePath())
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)
	return config, err
}

func (cm *ConfigManager) UpdateConfig(config model.AdrConfig) error {
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(PathResolverInst().ConfigFilePath(), bytes, 0644)
}

// Init initializes the ADR configuration folder
func (m *ConfigManager) InitConfig() error {

	baseDir := PathResolverInst().ConfigFolderPath()

	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err := os.Mkdir(baseDir, 0744); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Printf("Created directory %s\n", baseDir)
	} else {
		return fmt.Errorf("directory %s already exists. Not overriding", baseDir)
	}

	config := model.AdrConfig{
		BaseDir:    baseDir,
		CurrentAdr: 0,
	}
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}

	if err := os.WriteFile(PathResolverInst().ConfigFilePath(), bytes, 0644); err != nil {
		return fmt.Errorf("failed to write configuration: %w", err)
	}
	return nil
}
