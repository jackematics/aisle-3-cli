package main

import (
	"aisle-3-cli/command"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "list",
				Aliases: []string{"l"},
				Usage: "list pending jobs",
				Action: command.List,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}