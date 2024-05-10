package main

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Globals
	Init InitCmd `cmd:"" help:"Setup new ADR configuration in the current project"`
	New  NewCmd  `cmd:"" help:"Adds new ADR"`
	List ListCmd `cmd:"" help:"Lists all existing ADRs"`
}

func main() {

	cli := CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("adr"),
		kong.Description("ADR tool for your project"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
		kong.Vars{
			"version": "0.0.1",
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
