package cli

import (
	"errors"
	"testing"
)

func TestParseGitStatus(t *testing.T) {
	output := "?? .adr/001-new.md\n M .adr/002-existing.md\nR  .adr/003-old.md -> .adr/003-new.md\n"
	entries := parseGitStatus(output)

	if len(entries) != 3 {
		t.Fatalf("got %d entries, want 3", len(entries))
	}

	if entries[0].code != "??" || entries[0].path != ".adr/001-new.md" {
		t.Fatalf("unexpected first entry: %+v", entries[0])
	}

	if entries[2].path != ".adr/003-new.md" {
		t.Fatalf("expected rename target path, got %+v", entries[2])
	}
}

func TestDefaultADRCommitMessage(t *testing.T) {
	tests := []struct {
		name    string
		entries []gitStatusEntry
		want    string
	}{
		{
			name:    "single added ADR",
			entries: []gitStatusEntry{{code: "??", path: ".adr/001-new.md"}},
			want:    "Add ADR 001-new",
		},
		{
			name:    "single deleted ADR",
			entries: []gitStatusEntry{{code: "D", path: ".adr/002-old.md"}},
			want:    "Remove ADR 002-old",
		},
		{
			name:    "mixed ADR updates",
			entries: []gitStatusEntry{{code: "M", path: ".adr/001-one.md"}, {code: "??", path: ".adr/002-two.md"}},
			want:    "Update ADRs: 001-one, 002-two",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultADRCommitMessage(tt.entries)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCommitFailsOutsideGitRepo(t *testing.T) {
	oldIsGitRepo := isGitRepo
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
	})

	isGitRepo = func() bool { return false }

	err := (&CommitCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail")
	}
}

func TestCommitRefusesNonADRStagedChanges(t *testing.T) {
	oldIsGitRepo := isGitRepo
	oldRunGit := runGit
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
		runGit = oldRunGit
	})

	isGitRepo = func() bool { return true }
	runGit = func(args ...string) (string, error) {
		if len(args) >= 3 && args[0] == "diff" && args[1] == "--cached" {
			return "README.md", nil
		}
		t.Fatalf("unexpected git call: %v", args)
		return "", nil
	}

	err := (&CommitCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail")
	}
}

func TestCommitFailsWithoutADRChanges(t *testing.T) {
	oldIsGitRepo := isGitRepo
	oldRunGit := runGit
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
		runGit = oldRunGit
	})

	isGitRepo = func() bool { return true }
	runGit = func(args ...string) (string, error) {
		switch {
		case len(args) >= 3 && args[0] == "diff" && args[1] == "--cached":
			return "", nil
		case len(args) >= 4 && args[0] == "status" && args[1] == "--porcelain":
			return "", nil
		default:
			t.Fatalf("unexpected git call: %v", args)
			return "", nil
		}
	}

	err := (&CommitCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail")
	}
}

func TestCommitUsesGeneratedMessage(t *testing.T) {
	oldIsGitRepo := isGitRepo
	oldRunGit := runGit
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
		runGit = oldRunGit
	})

	isGitRepo = func() bool { return true }
	var commitArgs []string
	runGit = func(args ...string) (string, error) {
		switch {
		case len(args) >= 3 && args[0] == "diff" && args[1] == "--cached":
			return "", nil
		case len(args) >= 4 && args[0] == "status" && args[1] == "--porcelain":
			return "?? .adr/005-new.md", nil
		case len(args) == 3 && args[0] == "add":
			return "", nil
		case len(args) == 3 && args[0] == "commit":
			commitArgs = append([]string{}, args...)
			return "", nil
		default:
			t.Fatalf("unexpected git call: %v", args)
			return "", nil
		}
	}

	if err := (&CommitCmd{}).Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if len(commitArgs) != 3 || commitArgs[2] != "Add ADR 005-new" {
		t.Fatalf("unexpected commit args: %v", commitArgs)
	}
}

func TestCommitUsesExplicitMessage(t *testing.T) {
	oldIsGitRepo := isGitRepo
	oldRunGit := runGit
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
		runGit = oldRunGit
	})

	isGitRepo = func() bool { return true }
	var commitArgs []string
	runGit = func(args ...string) (string, error) {
		switch {
		case len(args) >= 3 && args[0] == "diff" && args[1] == "--cached":
			return "", nil
		case len(args) >= 4 && args[0] == "status" && args[1] == "--porcelain":
			return " M .adr/005-new.md", nil
		case len(args) == 3 && args[0] == "add":
			return "", nil
		case len(args) == 3 && args[0] == "commit":
			commitArgs = append([]string{}, args...)
			return "", nil
		default:
			t.Fatalf("unexpected git call: %v", args)
			return "", nil
		}
	}

	if err := (&CommitCmd{Message: "Custom ADR commit"}).Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if len(commitArgs) != 3 || commitArgs[2] != "Custom ADR commit" {
		t.Fatalf("unexpected commit args: %v", commitArgs)
	}
}

func TestCommitPropagatesGitError(t *testing.T) {
	oldIsGitRepo := isGitRepo
	oldRunGit := runGit
	t.Cleanup(func() {
		isGitRepo = oldIsGitRepo
		runGit = oldRunGit
	})

	isGitRepo = func() bool { return true }
	runGit = func(args ...string) (string, error) {
		if len(args) >= 3 && args[0] == "diff" && args[1] == "--cached" {
			return "", errors.New("git failed")
		}
		t.Fatalf("unexpected git call: %v", args)
		return "", nil
	}

	err := (&CommitCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail")
	}
}
