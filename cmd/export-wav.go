package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/timiskhakov/music/guitar"
	"github.com/timiskhakov/music/karplusstrong"
	"github.com/urfave/cli"
)

var ExportWav = cli.Command{
	Name:        "wav",
	Usage:       "ug w -id {tabId}",
	Description: "Fetch a tab by id and export it as a wav file",
	Aliases:     []string{"w"},
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "id",
			Value: 1947141,
			Usage: "",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "output.wav",
			Usage: "",
		},
	},
	Action: fetchTabByIDExportWav,
}

const sampleRate = 44100

func fetchTabByIDExportWav(c *cli.Context) {
	var tabID int64
	var outputFile string

	if c.IsSet("id") {
		tabID = c.Int64("id")
	}

	if c.IsSet("output") {
		outputFile = c.String("output")
	}

	s := ultimateguitar.New()
	tab, err := s.GetTabByID(tabID)

	if err != nil {
		log.Fatal(err)
	}

	content := tab.Content

	// find regex matches in content
	r, _ := regexp.Compile(`\[tab\](.|\n)*?\[\/tab]`)
	var chunks [][]string

	if r.MatchString(content) {
		matches := r.FindAllString(content, -1)
		for _, match := range matches {
			var chunk []string
			// fmt.Println("----")
			// fmt.Println(match)
			for _, line := range strings.Split(match, "\n") {
				x := line
				x = strings.ReplaceAll(x, "[tab]", "")
				x = strings.ReplaceAll(x, "[/tab]", "")
				if strings.Contains(x, "|") {
					chunk = append(chunk, x)
				}
			}
			if len(chunk) == 6 {
				chunks = append(chunks, chunk)
			}
		}
	} else {
		fmt.Println("Unable to parse tab. Debug:")
		fmt.Println(content)
		return
	}

	tLines := make([]string, 6)
	for cidx, chunk := range chunks {
		for idx, line := range chunk {
			c := strings.ReplaceAll(line[2:], "\n", "")
			if cidx == 0 {
				c = strings.ReplaceAll(line, "\n", "")
			}
			c = strings.ReplaceAll(c, "\r", "")
			tLines[idx] = fmt.Sprintf("%s%s", tLines[idx], c)
		}
	}

	tabContent := strings.Join(tLines, "\n")

	lines := strings.Split(tabContent, "\n")
	xCount := len(lines[0])

	type Notes struct {
		Time  int32
		Notes []guitar.Note
	}

	timings := make([]Notes, xCount)
	for xidx := range lines[0] {
		for stringNum, line := range lines {
			charVal := string(line[xidx])
			intVal, err := strconv.Atoi(charVal)
			if err == nil || intVal < 0 {
				note := guitar.Note{Fret: intVal, String: stringNum}
				timings[xidx].Time = int32(xidx)
				timings[xidx].Notes = append(timings[xidx].Notes, note)
			} else {
				timings[xidx].Time = int32(xidx)
			}
		}
	}

	// fmt.Println(timings)
	var blankTiming float64 = .01
	var noteTiming float64 = .3

	kse1 := karplusstrong.NewExtended(sampleRate, 0.5)
	g1 := guitar.NewGuitar(sampleRate, kse1)

	var sequence []beep.Streamer

	for _, timing := range timings {
		if len(timing.Notes) == 0 {
			sequence = append(sequence, g1.Silence(blankTiming)) // fill with silence!
		} else if len(timing.Notes) == 1 {
			sequence = append(sequence, g1.Pluck(guitar.Note{Fret: timing.Notes[0].Fret, String: timing.Notes[0].String + 1}, noteTiming))
		} else {
			var notes []guitar.Note
			for _, note := range timing.Notes {
				notes = append(notes, guitar.Note{Fret: note.Fret, String: note.String + 1})
			}
			sequence = append(sequence, g1.Chord(notes, noteTiming, 0.01))
		}
	}

	song := beep.Seq(sequence...)

	name := outputFile

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saving to %s...\n", name)
	if err = wav.Encode(f, song, beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3}); err != nil {
		panic(err)
	}

	fmt.Println("Saved.")
}
