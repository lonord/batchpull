package main

import "github.com/subchen/go-cli/v3"

var (
	depFlag    = &cli.Flag{Name: "d, depth", Usage: "Specify depth of directories to search", DefValue: "1"}
	depFlagVal = func(c *cli.Context) int { return c.GetInt("depth") }
)

var (
	cliFlags = []*cli.Flag{depFlag}
)
