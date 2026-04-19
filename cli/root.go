package cli

type Globals struct {
	Version VersionFlag `name:"version" short:"v" help:"Print version information and quit"`
}

type CLI struct {
	Globals
	Init       InitCmd       `cmd:"" help:"Setup ADR directory in the current project"`
	New        NewCmd        `cmd:"" help:"Creates new ADR using template"`
	List       ListCmd       `cmd:"" help:"Lists all existing ADRs"`
	Show       ShowCmd       `cmd:"" help:"Shows one ADR by number or slug"`
	Edit       EditCmd       `cmd:"" help:"Opens one ADR in an editor"`
	Last       LastCmd       `cmd:"" help:"Shows the newest ADR"`
	DropLast   DropLastCmd   `cmd:"" name:"drop-last" help:"Deletes the newest ADR if it is not in a final state"`
	Commit     CommitCmd     `cmd:"" help:"Stages and commits ADR changes only"`
	Completion CompletionCmd `cmd:"" help:"Generate shell completion scripts"`
}
