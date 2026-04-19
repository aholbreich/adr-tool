package cli

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/aholbreich/adr-tool/internal/config"
)

var runGit = func(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	trimmed := strings.TrimSpace(string(output))
	if err != nil {
		if trimmed == "" {
			return "", err
		}
		return "", fmt.Errorf("%w: %s", err, trimmed)
	}
	return trimmed, nil
}

var isGitRepo = func() bool {
	return config.PathResolverInst().IsFilepathGitRepo()
}

// CLI Command
type CommitCmd struct {
	Message string `name:"message" short:"m" help:"Commit message to use for ADR changes"`
}

// Command Handler
func (r *CommitCmd) Run() error {
	if !isGitRepo() {
		return fmt.Errorf("current directory is not a git repository")
	}

	stagedFiles, err := runGit("diff", "--cached", "--name-only")
	if err != nil {
		return fmt.Errorf("inspect staged changes: %w", err)
	}

	for _, path := range splitLines(stagedFiles) {
		if !strings.HasPrefix(path, ".adr/") {
			return fmt.Errorf("refusing to commit because non-ADR staged change exists: %s", path)
		}
	}

	statusOutput, err := runGit("status", "--porcelain", "--", ".adr")
	if err != nil {
		return fmt.Errorf("inspect ADR changes: %w", err)
	}

	entries := parseGitStatus(statusOutput)
	if len(entries) == 0 {
		return fmt.Errorf("no ADR changes to commit")
	}

	if _, err := runGit("add", "--", ".adr"); err != nil {
		return fmt.Errorf("stage ADR changes: %w", err)
	}

	message := strings.TrimSpace(r.Message)
	if message == "" {
		message = defaultADRCommitMessage(entries)
	}

	if _, err := runGit("commit", "-m", message); err != nil {
		return fmt.Errorf("commit ADR changes: %w", err)
	}

	fmt.Printf("Committed ADR changes with message: %s\n", message)
	return nil
}

type gitStatusEntry struct {
	code string
	path string
}

func parseGitStatus(output string) []gitStatusEntry {
	lines := splitLines(output)
	entries := make([]gitStatusEntry, 0, len(lines))
	for _, line := range lines {
		if len(line) < 4 {
			continue
		}

		path := strings.TrimSpace(line[3:])
		if path == "" {
			continue
		}
		if parts := strings.Split(path, " -> "); len(parts) == 2 {
			path = strings.TrimSpace(parts[1])
		}

		entries = append(entries, gitStatusEntry{
			code: strings.TrimSpace(line[:2]),
			path: path,
		})
	}
	return entries
}

func defaultADRCommitMessage(entries []gitStatusEntry) string {
	action := "Update"
	allAdded := true
	allDeleted := true
	titles := make([]string, 0, len(entries))

	for _, entry := range entries {
		switch entry.code {
		case "A", "??":
			allDeleted = false
		case "D":
			allAdded = false
		default:
			allAdded = false
			allDeleted = false
		}

		titles = append(titles, strings.TrimSuffix(strings.TrimPrefix(entry.path, ".adr/"), ".md"))
	}

	if allAdded {
		action = "Add"
	} else if allDeleted {
		action = "Remove"
	}

	if len(titles) == 1 {
		return fmt.Sprintf("%s ADR %s", action, titles[0])
	}

	return fmt.Sprintf("%s ADRs: %s", action, strings.Join(titles, ", "))
}

func splitLines(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return strings.Split(strings.TrimSpace(s), "\n")
}
