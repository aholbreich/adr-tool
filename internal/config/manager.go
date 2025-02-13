package config

import (
	"adr-tool/internal/model"
	"encoding/json"
	"os"
)

const configFilePath = ".adr/config.json"

type ConfigManager struct{}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

func (cm *ConfigManager) LoadConfig() (model.AdrConfig, error) {
	var config model.AdrConfig
	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)
	return config, err
}

func (cm *ConfigManager) SaveConfig(config model.AdrConfig) error {
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, bytes, 0644)
}
