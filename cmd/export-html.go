package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"

	"github.com/urfave/cli"
)

//go:embed data/templates/template.tmpl
var templateContent string

var ExportTabHTML = cli.Command{
	Name:        "export",
	Usage:       "ug export -id {tabId}",
	Description: "Fetch a tab from ultimate-guitar.com by id and print it as HTML",
	Aliases:     []string{"e"},
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "id",
			Value: 1086983,
			Usage: "",
		},
	},
	Action: exportTabByID,
}

func exportTabByID(c *cli.Context) {
	var tabID int64

	if c.IsSet("id") {
		tabID = c.Int64("id")
	}

	s := ultimateguitar.New()
	tab, err := s.GetTabByID(tabID)

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("tab").Parse(templateContent)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer

	err = t.Execute(&out, tab)
	if err != nil {
		log.Fatal(err)
	}

	// hood shit ayyy
	html := strings.ReplaceAll(out.String(), "[tab]", "<div class=\"gtab\">")
	html = strings.ReplaceAll(html, "[/tab]", "</div>")
	html = strings.ReplaceAll(html, "[ch]", "<span class=\"chord\">")
	html = strings.ReplaceAll(html, "[/ch]", "</span>")

	fmt.Println(html)
}
