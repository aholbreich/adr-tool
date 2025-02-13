package cli

type Globals struct {
	Version VersionFlag `name:"version" short:"v" help:"Print version information and quit"`
}

type CLI struct {
	Globals
	Init InitCmd `cmd:"" help:"Setup new ADR configuration in the current project"`
	New  NewCmd  `cmd:"" help:"Adds new ADR"`
	List ListCmd `cmd:"" help:"Lists all existing ADRs"`
}
