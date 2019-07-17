package main

import (
	"fmt"
	"log"

	"github.com/Pilfer/ultimate-guitar-scraper/scraper"
)

func main() {
	fmt.Println("Running..")
	s := scraper.New()
	// s.SetProxy("http://localhost:8080")
	fmt.Println("Scraper intialized with device id of: ", s.DeviceID)
	tab, err := s.GetTabByID(1956589)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Song name:", tab.SongName, " by ", tab.ArtistName)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println(tab.Content)
	fmt.Println("----------------------------------------------------------------------")

}
