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
			fmt.Println("Initialization aborted by the user.")
			return nil
		}
	}

	mgr := config.NewConfigManager()
	if err := mgr.InitConfig(); err != nil {
		return fmt.Errorf("initialize ADRs: %w", err)
	}

	fmt.Println("ADR initialized successfully at", pathResolver.ConfigFolderPath())
	return nil
}
