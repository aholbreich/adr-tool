package cli

import (
	"fmt"

	"github.com/aholbreich/adr-tool/internal/adr"
	"github.com/aholbreich/adr-tool/internal/config"
)

// CLI Command
type ListCmd struct{}

// Command Handler
func (r *ListCmd) Run() error {

	adrs, err := adr.NewAdrManager().GetADRList()
	if err != nil {
		return fmt.Errorf("failed to list ADRs: %w", err)
	}

	if len(adrs) == 0 {
		pathResolver := config.PathResolverInst()
		fmt.Printf("No ADRs found in %s.\n", pathResolver.ConfigFolderPath())
		return nil
	}

	fmt.Println("Architecture Decision Records:")
	for _, adr := range adrs {
		fmt.Println(" -", adr)
	}

	return nil
}
