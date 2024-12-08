// cmd/count_chords.go
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

// CountChords is the CLI command to count chord occurrences across all songs
var CountChords = cli.Command{
	Name:        "count_chords",
	Usage:       "Counts the appearance and frequency of chords in all songs",
	Description: "Analyzes all song files in the specified output directory and generates a statistics report of chord usage.",
	Aliases:     []string{"cc"},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "input",
			Usage: "--input {input directory}. Default './out'",
		},
		cli.StringFlag{
			Name:  "output",
			Usage: "--output {output file path}. Default './chord_stats.txt'",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug logging",
		},
	},
	Action: CountChordsAction,
}

// CountChordsAction is the action function for the count_chords command
func CountChordsAction(c *cli.Context) {
	if c.Bool("debug") {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	inputDir := "./out"
	if c.IsSet("input") {
		inputDir = c.String("input")
	}

	outputFile := "./chord_stats.txt"
	if c.IsSet("output") {
		outputFile = c.String("output")
	}

	chordCounts := make(map[string]int)

	// Improved regular expression to match a wide range of chords
	// Matches chords like C, Cm, Cmaj7, Cadd9, C#, Db, Bm/D, etc.
	chordRegex := regexp.MustCompile(`\b[A-G](?:#|b)?(?:m|maj7|add9|sus4|dim|aug|7|9|11|13|maj|m7|m9|aug7|dim7)?(?:/[A-G](?:#|b)?)?\b`)

	// Regular expression to detect tablature lines
	tablatureRegex := regexp.MustCompile(`^[eBgdAE]\|`)

	// Iterate over all files in the input directory
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process .crd files (assuming song files have .crd extension)
		if strings.ToLower(filepath.Ext(path)) != ".crd" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Printf("Error opening file %s: %v", path, err)
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 0
		for scanner.Scan() {
			line := scanner.Text()
			lineNumber++

			// Skip metadata lines
			if strings.HasPrefix(line, "{") || strings.HasPrefix(line, "[") {
				continue
			}

			// Skip tablature lines
			if tablatureRegex.MatchString(line) {
				continue
			}

			// Find all chord matches in the line
			matches := chordRegex.FindAllString(line, -1)
			for _, chord := range matches {
				normalizedChord := normalizeChord(chord)
				if normalizedChord != "" {
					chordCounts[normalizedChord]++
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through input directory: %v", err)
	}

	// Sort chords by frequency
	type chordFrequency struct {
		Chord string
		Count int
	}

	var frequencies []chordFrequency
	for chord, count := range chordCounts {
		frequencies = append(frequencies, chordFrequency{Chord: chord, Count: count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	// Prepare the output
	outputLines := []string{
		"Chord Usage Statistics",
		"======================",
	}
	for _, freq := range frequencies {
		line := fmt.Sprintf("%s: %d", freq.Chord, freq.Count)
		outputLines = append(outputLines, line)
	}

	// Write to the output file
	err = os.WriteFile(outputFile, []byte(strings.Join(outputLines, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to output file %s: %v", outputFile, err)
	}

	fmt.Printf("Chord statistics written to %s\n", outputFile)
}

// normalizeChord standardizes chord notation for consistent counting
func normalizeChord(chord string) string {
	// Remove any surrounding whitespace and convert to proper case
	chord = strings.TrimSpace(chord)
	if chord == "" {
		return ""
	}

	// Split chord and inversion if present
	var chordPart, inversionPart string
	if strings.Contains(chord, "/") {
		parts := strings.Split(chord, "/")
		chordPart = parts[0]
		inversionPart = parts[1]
	} else {
		chordPart = chord
	}

	// Normalize chord part
	chordPart = normalizeChordPart(chordPart)

	// Normalize inversion part if present
	if inversionPart != "" {
		inversionPart = normalizeChordPart(inversionPart)
		return chordPart + "/" + inversionPart
	}
	return chordPart
}

// normalizeChordPart formats the chord part correctly
func normalizeChordPart(chord string) string {
	if len(chord) == 0 {
		return ""
	}

	// Capitalize the first letter and handle the rest
	chord = strings.ToUpper(string(chord[0])) + chord[1:]

	// Convert minor indicators to lowercase 'm'
	chord = strings.ReplaceAll(chord, "M", "m")
	// Handle specific chord suffixes
	suffixes := []string{"Maj", "Maj7", "m7", "m9", "m11", "m13", "add9", "sus4", "dim", "dim7", "aug", "aug7", "7", "9", "11", "13"}

	for _, suffix := range suffixes {
		if strings.HasSuffix(chord, suffix) {
			index := strings.LastIndex(chord, suffix)
			chord = chord[:index] + suffix
			break
		}
	}

	return chord
}
