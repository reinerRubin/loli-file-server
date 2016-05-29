package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/reinerRubin/loli-file-server"
)

func main() {
	app := cli.NewApp()
	app.Name = "Loli-server client"

	app.Commands = []cli.Command{
		{
			Name:      "ls",
			Usage:     "address",
			UsageText: "./client ls localhost:50051",
			Action: func(c *cli.Context) error {
				return loli.NewFilerCli(c).PrintFileList()
			},
		},
		{
			Name:      "cp",
			Usage:     "cp address file1 fileN dst",
			UsageText: "./client cp localhost:50051 file /tmp/file",
			Action: func(c *cli.Context) error {
				return loli.NewFilerCli(c).SaveFiles()
			},
		},
	}

	app.Run(os.Args)
}
