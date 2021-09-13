package main

import (
	"os"

	"github.com/subchen/go-cli/v3"
)

func main() {
	app := cli.NewApp()
	app.Name = "batchpull"
	app.Version = "0.1.1"
	app.Usage = "A tool for batch updating git repositories"
	app.Flags = cliFlags
	app.Action = runApp
	app.Run(os.Args)
}
