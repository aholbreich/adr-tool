package cli

import (
	"fmt"
	"strings"

	"adr-tool/internal/adr"
	"adr-tool/internal/config"
)

// CLI Command
type NewCmd struct {
	AdrName []string `arg:"" required:"" help:"ADR Name"`
}

// Command Handler
func (r *NewCmd) Run() error {
	adrName := strings.Join(r.AdrName, " ")

	// Load configuration using ConfigManager
	cfgManager := config.NewConfigManager()
	currentConfig, err := cfgManager.LoadConfig()
	if err != nil {
		fmt.Println("No ADR configuration found!")
		fmt.Println("Start by initializing ADR configuration, check 'adr init --help' for more help")
		return err
	}

	// Increment ADR number and update config
	currentConfig.CurrentAdr++
	if err := cfgManager.SaveConfig(currentConfig); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	// Create the new ADR

	adrManager := adr.NewManager(config.NewPathResolver())

	if err := adrManager.CreateNewAdr(currentConfig, adrName); err != nil {
		return fmt.Errorf("failed to create new ADR: %w", err)
	}

	fmt.Printf("New ADR %03d was successfully written to: %s\n", currentConfig.CurrentAdr, currentConfig.BaseDir)
	return nil
}
