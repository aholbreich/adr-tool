package cli

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"adr-tool/model"
)

// CLI Command
type NewCmd struct {
	AdrName []string `arg:"" required:"" help:"ADR Name"`
}

// Command Handler
func (r *NewCmd) Run() error {
	adrName := strings.Join(r.AdrName, " ")

	// Load existing configuration
	currentConfig, err := getConfig()
	if err != nil {
		fmt.Println("No ADR configuration found!")
		fmt.Println("Start by initializing ADR configuration, check 'adr init --help' for more help")
		return err
	}

	// Increment ADR number and update config
	currentConfig.CurrentAdr++
	if err := updateConfig(currentConfig); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	// Create the new ADR
	if err := newAdr(currentConfig, adrName); err != nil {
		return fmt.Errorf("failed to create new ADR: %w", err)
	}

	fmt.Printf("New ADR %d was successfully written to: %s\n", currentConfig.CurrentAdr, configFolderPath)
	return nil
}

// Load ADR configuration from file
func getConfig() (model.AdrConfig, error) {
	var currentConfig model.AdrConfig

	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return currentConfig, err
	}

	if err := json.Unmarshal(bytes, &currentConfig); err != nil {
		return currentConfig, err
	}

	return currentConfig, nil
}

// Save updated ADR configuration to file
func updateConfig(config model.AdrConfig) error {
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, bytes, 0644)
}

// Create a new ADR file with the given name
func newAdr(config model.AdrConfig, adrName string) error {
	adr := model.Adr{
		Title:  adrName,
		Date:   time.Now().Format("2006-01-02 15:04"), // ISO 8601 format
		Number: config.CurrentAdr,
		Status: model.PROPOSED,
	}

	tpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Sanitize and build the filename
	adrFileName := fmt.Sprintf("%03d-%s.md", adr.Number, strings.Join(strings.Split(strings.TrimSpace(adr.Title), " "), "-"))
	adrFullPath := filepath.Join(config.BaseDir, adrFileName)

	f, err := os.Create(adrFullPath)
	if err != nil {
		return fmt.Errorf("failed to create ADR file: %w", err)
	}
	defer f.Close()

	// Render template into the new ADR file
	if err := tpl.Execute(f, adr); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
