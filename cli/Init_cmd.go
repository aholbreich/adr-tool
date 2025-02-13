package cli

import (
	"adr-tool/internal/adr"
	"adr-tool/internal/config"
	"fmt"
)

// CLI Command
type InitCmd struct {
}

// Command Handler
func (r *InitCmd) Run() error {

	pathResolver := config.NewPathResolver()

	if !pathResolver.IsFilepathGitRepo() {
		if !confirmAction(".git folder is not detected. This does not seem to be the root of your project. Do you still want to proceed (Y/n)?") {
			fmt.Printf("Initialization aborted by the user.\n")
			return nil
		}
	}

	//TODO should be Config Manager
	adrManager := adr.NewManager(pathResolver)

	if err := adrManager.InitConfig(); err != nil {
		fmt.Println("Failed to initialize ADRs:", err)
		return nil
	}

	fmt.Println("ADR initialized successfully at", pathResolver.ConfigFolderPath())
	return nil
}
