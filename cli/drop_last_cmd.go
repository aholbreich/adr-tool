package cli

import (
	"fmt"
	"os"

	"github.com/aholbreich/adr-tool/internal/adr"
	"github.com/aholbreich/adr-tool/internal/model"
)

var listADRs = func() ([]model.ADR, error) {
	return adr.NewADRManager().ListADRs()
}

var lastADRPath = func() (string, error) {
	return adr.NewADRManager().LastADR()
}

var removeADRFile = os.Remove
var confirmDropLast = confirmAction

// CLI Command
type DropLastCmd struct {
	Yes bool `name:"yes" help:"Delete without confirmation"`
}

// Command Handler
func (r *DropLastCmd) Run() error {
	adrs, err := listADRs()
	if err != nil {
		return fmt.Errorf("list ADRs: %w", err)
	}

	if len(adrs) == 0 {
		return fmt.Errorf("no ADRs found")
	}

	last := adrs[0]
	if isFinalADRStatus(last.Status) {
		return fmt.Errorf("refusing to delete ADR %s because status is %s", last.Title, last.Status)
	}

	adrPath, err := lastADRPath()
	if err != nil {
		return fmt.Errorf("find last ADR: %w", err)
	}

	if !r.Yes {
		prompt := fmt.Sprintf("Delete %s [%s] (Y/n)? ", last.Title, last.Status)
		if !confirmDropLast(prompt) {
			fmt.Println("Deletion aborted by the user.")
			return nil
		}
	}

	if err := removeADRFile(adrPath); err != nil {
		return fmt.Errorf("delete ADR: %w", err)
	}

	fmt.Printf("Deleted ADR %s\n", last.Title)
	return nil
}

func isFinalADRStatus(status model.ADRStatus) bool {
	switch status {
	case model.StatusAccepted, model.StatusDeprecated, model.StatusSuperseded:
		return true
	default:
		return false
	}
}
