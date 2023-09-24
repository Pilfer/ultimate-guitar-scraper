package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"

	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
	"github.com/urfave/cli"
	"golang.org/x/term"
)

var GetAll = cli.Command {
	Name:	"get_all",
	Usage:	"ug a",
	Description: "Fetches all saved tabs/songs for ultimate-guitar. Requires you to login",
	Aliases:	[]string{"a"},
	Flags:	[]cli.Flag{
		cli.StringFlag{
			Name: "user",
			Usage: "--user {your_email} ",
		},
		cli.StringFlag{
			Name: "output",
			Usage: "--output {output path}. Default './out'",
		},
	},
	Action: GetAllTabs,
}
func GetAllTabs(c *cli.Context) {
	var user, password string
	var err error

	if c.IsSet("user") {
		user = c.String("user")
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Username: ")
		user, err = reader.ReadString('\n')
		user = strings.Trim(user, "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password = string(bytePassword)

	tabs, err := fetchAllTabs(user, password)
	if err != nil {
		log.Fatal(err)
	}

	path := "./out/"
	if c.IsSet("output") {
		path = c.String("output")
	}

	path, err = filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	err = writeTabs(path, tabs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d tabs to %s", len(tabs), path)

}
func fetchAllTabs(user string, password string) ([]ultimateguitar.TabResult, error) {
	var tabResults []ultimateguitar.TabResult
	s := ultimateguitar.New()
	res, err := s.Login(user, password)
	if err != nil || res == "Failed to login"{
		return tabResults, errors.New("Failed to login")
	}
	tabResults, err = s.GetAll()
	return tabResults, err
}

func writeTabs(path string, tabs []ultimateguitar.TabResult) error {
	if path == "" {
		return errors.New("writeTabs: requires path")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0775)
		if err != nil {
			return err
		}
	}
	fmt.Println(path)

	for i := 0; i < len(tabs); i++ {
		artist := tabs[i].ArtistName
		songName := tabs[i].SongName
		content := tabs[i].Content
		content = fmt.Sprintf("{artist: %s}\n{title: %s}\n%s", artist, songName, content)

		regex := regexp.MustCompile(`\[(/?tab|/?ch)\]`)
		content = regex.ReplaceAllString(content, "")

		filename := fmt.Sprintf("%s-%s.crd", artist, songName)
		err := os.WriteFile(fmt.Sprintf("%s/%s", path, filename), []byte(content), 0644)
		if err != nil {
			log.Print(fmt.Sprintf("Error writing file %s", filename))
			log.Print(err)
			continue 
		}
	}
	return nil
}