package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/aholbreich/adr-tool/internal/adr"
)

var lookupEnv = os.LookupEnv
var lookPath = exec.LookPath
var currentGOOS = runtime.GOOS
var runEditor = func(editor, filePath string) error {
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CLI Command
type EditCmd struct {
	ID string `arg:"" required:"" help:"ADR number, full stem, or slug"`
}

// Command Handler
func (r *EditCmd) Run() error {
	adrPath, err := adr.NewADRManager().FindADR(r.ID)
	if err != nil {
		return fmt.Errorf("find ADR: %w", err)
	}

	editor, err := resolveEditor()
	if err != nil {
		return err
	}

	if err := runEditor(editor, adrPath); err != nil {
		return fmt.Errorf("run editor %q: %w", editor, err)
	}

	return nil
}

func resolveEditor() (string, error) {
	candidates := []string{}
	for _, envName := range []string{"VISUAL", "EDITOR"} {
		if value, ok := lookupEnv(envName); ok && value != "" {
			candidates = append(candidates, value)
		}
	}

	if len(candidates) == 0 {
		candidates = defaultEditors()
	}

	for _, candidate := range candidates {
		if _, err := lookPath(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("no editor found; set $VISUAL or $EDITOR")
}

func defaultEditors() []string {
	if currentGOOS == "windows" {
		return []string{"notepad"}
	}

	return []string{"vim", "vi", "nano"}
}
