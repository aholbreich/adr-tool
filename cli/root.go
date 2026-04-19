package cli

type Globals struct {
	Version VersionFlag `name:"version" short:"v" help:"Print version information and quit"`
}

type CLI struct {
	Globals
	Init InitCmd `cmd:"" help:"Setup ADR directory in the current project"`
	New  NewCmd  `cmd:"" help:"Creates new ADR using template"`
	List ListCmd `cmd:"" help:"Lists all existing ADRs"`
	Show ShowCmd `cmd:"" help:"Shows one ADR by number or slug"`
}
