package main

import (
	"log"
	"os"

	"github.com/Pilfer/ultimate-guitar-scraper/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Before = cmd.BeforeCommand
	app.After = cmd.AfterCommand
	app.Name = "ug"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		cmd.FetchTab,
		cmd.ExportTabHTML,
		cmd.ExportWav,
		cmd.GetAll,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal("An error occurred: ", err)
	}
}
