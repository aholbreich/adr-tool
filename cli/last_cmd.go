package cli

import (
	"fmt"
	"os"

	"github.com/aholbreich/adr-tool/internal/adr"
)

// CLI Command
type LastCmd struct{}

// Command Handler
func (r *LastCmd) Run() error {
	adrPath, err := adr.NewADRManager().LastADR()
	if err != nil {
		return fmt.Errorf("find last ADR: %w", err)
	}

	content, err := os.ReadFile(adrPath)
	if err != nil {
		return fmt.Errorf("read ADR: %w", err)
	}

	fmt.Print(string(content))
	return nil
}
