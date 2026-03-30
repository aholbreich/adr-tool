package cli

import (
	"fmt"
	"strings"

	"github.com/aholbreich/adr-tool/internal/adr"
	"github.com/aholbreich/adr-tool/internal/config"
)

// CLI Command
type NewCmd struct {
	AdrName []string `arg:"" required:"" help:"ADR Name"`
}

// Command Handler
func (r *NewCmd) Run() error {
	adrName := strings.Join(r.AdrName, " ")

	cfgManager := config.NewConfigManager()
	currentConfig, err := cfgManager.LoadConfig()
	if err != nil {
		return fmt.Errorf("load ADR configuration: %w; run 'adr init' first", err)
	}

	currentConfig.CurrentAdr++
	if err := cfgManager.UpdateConfig(currentConfig); err != nil {
		return fmt.Errorf("update ADR config: %w", err)
	}

	if err := adr.NewAdrManager().CreateNewAdr(currentConfig, adrName); err != nil {
		return fmt.Errorf("create new ADR: %w", err)
	}

	fmt.Printf("New ADR %03d was successfully written to: %s\n", currentConfig.CurrentAdr, currentConfig.BaseDir)
	return nil
}
