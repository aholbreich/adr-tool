package config

import (
	"fmt"
	"os"
)

type ConfigManager struct{}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

// InitConfig initializes the ADR directory.
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

	return nil
}
