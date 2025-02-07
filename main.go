package main

import (
	"adr-tool/cli"

	"github.com/alecthomas/kong"
)

var version = "dev" // Default to "dev", overridden by build flags

func main() {

	cli := cli.CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("adr"),
		kong.Description("ADR tool for your project. \n Project details can be found at https://github.com/aholbreich/adr-tool"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
		kong.Vars{
			"version": version,
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
