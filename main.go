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
			{
				Name: "complete",
				Aliases: []string{"c"},
				Usage: "complete a job",
				Action: command.Complete,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "room",
						Aliases: []string{"r"},
						Usage: "room job was completed in",
					},
					&cli.StringSliceFlag{
						Name: "jobs",
						Aliases: []string{"j"},
						Usage: "jobs completed",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}