package cli

import (
	"fmt"
	"os"
	"sort"
	"unicode"
)

// CLI Command
type ListCmd struct{}

// Command Handler
func (r *ListCmd) Run() error {
	adrs, err := r.listADRs()
	if err != nil {
		return fmt.Errorf("failed to list ADRs: %w", err)
	}

	if len(adrs) == 0 {
		fmt.Printf("No ADRs found in %s.\n", configFolderPath)
		return nil
	}

	fmt.Println("Architecture Decision Records:")
	for _, adr := range adrs {
		fmt.Println(" -", adr)
	}

	return nil
}

// List ADR files in reverse order
func (r *ListCmd) listADRs() ([]string, error) {
	entries, err := os.ReadDir(configFolderPath)
	if err != nil {
		return nil, err
	}

	var adrs []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if len(name) > 0 && unicode.IsDigit(rune(name[0])) { // Starts with digit? Must be an ADR file
			adrs = append(adrs, name)
		}
	}

	// Reverse order by sorting in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(adrs)))
	return adrs, nil
}
