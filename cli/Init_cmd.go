package cli

import (
	"fmt"

	"github.com/aholbreich/adr-tool/internal/config"
)

// CLI Command
type InitCmd struct {
}

// Command Handler
func (r *InitCmd) Run() error {

	pathResolver := config.PathResolverInst()

	if !pathResolver.IsFilepathGitRepo() {
		if !confirmAction(".git folder is not detected. This does not seem to be the root of your project. Do you still want to proceed (Y/n)?") {
			fmt.Printf("Initialization aborted by the user.\n")
			return nil
		}
	}

	mgr := config.NewConfigManager()

	if err := mgr.InitConfig(); err != nil {
		fmt.Println("Failed to initialize ADRs:", err)
		return nil
	}

	fmt.Println("ADR initialized successfully at", pathResolver.ConfigFolderPath())
	return nil
}
