package cli

import (
	"fmt"
	"os"

	"github.com/aholbreich/adr-tool/internal/adr"
)

// CLI Command
type ShowCmd struct {
	ID string `arg:"" required:"" help:"ADR number, full stem, or slug"`
}

// Command Handler
func (r *ShowCmd) Run() error {
	adrPath, err := adr.NewADRManager().FindADR(r.ID)
	if err != nil {
		return fmt.Errorf("find ADR: %w", err)
	}

	content, err := os.ReadFile(adrPath)
	if err != nil {
		return fmt.Errorf("read ADR: %w", err)
	}

	fmt.Print(string(content))
	return nil
}
