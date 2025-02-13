package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"adr-tool/config"
	"adr-tool/model"

	"github.com/fatih/color"
)

type InitCmd struct {
}

// Command Handler
func (r *InitCmd) Run(ctx *Globals) error {
	color.White("Initializing ADR configuration at %s", configFolderPath)

	if !isGitRepo() {
		if !confirmAction(".git folder is not detected. This does not seem to be the root of your project. Do you still want to proceed (Y/n)?") {
			color.Red("Initialization aborted by the user.")
			return nil
		}
	}

	if err := initConfig(configFolderPath); err != nil {
		return fmt.Errorf("failed to initialize configuration: %w", err)
	}
	if err := initTemplate(); err != nil {
		return fmt.Errorf("failed to initialize template: %w", err)
	}

	color.Green("ADR configuration initialized successfully.")
	return nil
}

func isGitRepo() bool {
	if _, err := os.Stat(filepath.Join(".", ".git")); os.IsNotExist(err) {
		return false
	}
	return true
}

// Initialize ADR Configuration
func initConfig(baseDir string) error {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err := os.MkdirAll(baseDir, 0744); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	} else {
		color.Red("Directory %s already exists. Not overriding.", baseDir)
		os.Exit(-1)
	}
	config := model.AdrConfig{BaseDir: baseDir, CurrentAdr: 0}
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}
	color.White("Writing new configuration at: %s", configFilePath)
	if err := os.WriteFile(configFilePath, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write configuration: %w", err)
	}
	return nil
}

// Initialize Template
func initTemplate() error {
	body := []byte(config.TEMPLATE1)
	if err := os.WriteFile(templateFilePath, body, 0644); err != nil {
		return fmt.Errorf("failed to write template: %w", err)
	}
	return nil
}
