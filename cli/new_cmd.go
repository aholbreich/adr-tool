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
	baseDir := config.PathResolverInst().ConfigFolderPath()

	nextNumber, err := adr.NewADRManager().NextADRNumber(baseDir)
	if err != nil {
		return fmt.Errorf("determine next ADR number: %w; run 'adr init' first", err)
	}

	_, err = adr.NewADRManager().CreateNewADR(baseDir, nextNumber, adrName)
	if err != nil {
		return fmt.Errorf("create new ADR: %w", err)
	}

	fmt.Printf("New ADR %03d was successfully written to: %s\n", nextNumber, baseDir)
	return nil
}
