package cmd

import (
	"fmt"
	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
	"github.com/urfave/cli"
	"log"
	"strings"
)

var FetchTab = cli.Command{
	Name:        "fetch",
	Usage:       "ug f -id {tabId}",
	Description: "Fetch a tab from ultimate-guitar.com by id",
	Aliases:     []string{"f"},
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "id",
			Value: 1947141,
			Usage: "",
		},
	},
	Action: fetchTabByID,
}

func fetchTabByID(c *cli.Context) {
	var tabID int64

	if c.IsSet("id") {
		tabID = c.Int64("id")
	}

	s := ultimateguitar.New()
	tab, err := s.GetTabByID(tabID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Song name:", tab.SongName, " by ", tab.ArtistName)
	fmt.Println("----------------------------------------------------------------------")

	// Remove the syntax delimiters as a proof of concept
	tabOut := tab.Content
	tabOut = strings.ReplaceAll(tabOut, "[tab]", "")
	tabOut = strings.ReplaceAll(tabOut, "[/tab]", "")
	tabOut = strings.ReplaceAll(tabOut, "[ch]", "")
	tabOut = strings.ReplaceAll(tabOut, "[/ch]", "")
	fmt.Println(tabOut)
	fmt.Println("----------------------------------------------------------------------")
}
