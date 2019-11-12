package main

import (
	"fmt"
	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
)

func main() {
	fmt.Println("Running..")
	s := ultimateguitar.New()
	s.SetProxy("http://localhost:8888")
	fmt.Println("Scraper initialized with device id of: ", s.DeviceID)

	_, _ = s.TabByURL("https://tabs.ultimate-guitar.com/tab/johnny_flynn/raising_the_dead_chords_1947141")

	/*
	tab, err := s.GetTabByID(1956589)

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
	*/

}
