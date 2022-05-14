package main

import (
	"fmt"
	"log"

	"github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
)

func main() {

	s := ultimateguitar.New()

	applicatures, err := s.GetApplicature("guitar", "E A D G B E", []ultimateguitar.ChordKey{
		ultimateguitar.ChordKeyA,
		ultimateguitar.ChordKeyG,
		ultimateguitar.ChordKeyG6,
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, app := range applicatures {
		fmt.Println("\n----------------------------------------------------------------------")
		fmt.Printf("%s has %d variations\nBase chord:\t", app.Chord, len(app.Variations))
		for _, fret := range app.Variations[0].Frets {
			if fret < 0 {
				fmt.Printf("-x")
			} else {
				fmt.Printf("-%d", fret)
			}
		}

		// Loop over app.Variations but skip the first one
		for vidx, variation := range app.Variations[1:] {
			fmt.Printf("\n%s Variant #%d:\t", app.Chord, vidx)
			for _, fret := range variation.Frets {
				if fret < 0 {
					fmt.Printf("-x")
				} else {
					fmt.Printf("-%d", fret)
				}
			}
		}
	}

	fmt.Println("\nAll done!")
}

/*
Result:

----------------------------------------------------------------------
A has 23 variations
Base chord:     -0-2-2-2-0-x
A Variant #0:   -5-2-2-2-0-x
A Variant #1:   -5-5-6-7-0-x
A Variant #2:   -5-5-6-7-7-5
A Variant #3:   -5-5-6-7-x-x
A Variant #4:   -9-10-9-7-0-x
A Variant #5:   -9-10-9-7-x-x
A Variant #6:   -9-10-9-11-0-x
A Variant #7:   -12-10-9-11-0-x
A Variant #8:   -9-10-9-11-12-x
A Variant #9:   -x-2-2-2-0-x
A Variant #10:  -x-2-2-2-4-5
A Variant #11:  -x-5-6-7-0-x
A Variant #12:  -x-5-6-7-7-5
A Variant #13:  -x-10-9-11-0-x
A Variant #14:  -x-10-9-11-12-x
A Variant #15:  -x-x-2-2-4-5
A Variant #16:  -x-x-6-7-7-5
A Variant #17:  -0-2-2-x-0-x
A Variant #18:  -5-5-6-7-x-5
A Variant #19:  -9-10-9-x-12-x
A Variant #20:  -x-2-2-2-x-5
A Variant #21:  -x-5-6-7-x-5
----------------------------------------------------------------------
G6 has 42 variations
Base chord:     -0-0-0-0-2-3
G6 Variant #0:  -0-3-0-0-2-3
G6 Variant #1:  -0-0-4-0-2-3
G6 Variant #2:  -0-0-0-2-2-3
G6 Variant #3:  -0-3-4-0-2-3
G6 Variant #4:  -3-0-0-2-2-3
G6 Variant #5:  -0-3-0-2-2-3
G6 Variant #6:  -0-0-4-2-2-3
G6 Variant #7:  -0-3-4-5-x-x
G6 Variant #8:  -3-3-4-5-7-3
G6 Variant #9:  -3-5-4-5-x-x
G6 Variant #10: -7-5-7-5-x-x
G6 Variant #11: -12-12-9-9-10-x
G6 Variant #12: -12-12-12-9-10-x
G6 Variant #13: -12-12-12-12-10-x
G6 Variant #14: -x-0-0-2-2-3
G6 Variant #15: -x-3-0-2-2-3
G6 Variant #16: -x-0-4-2-2-3
G6 Variant #17: -x-3-4-2-2-3
G6 Variant #18: -x-5-4-2-2-3
G6 Variant #19: -x-8-9-9-10-x
G6 Variant #20: -x-12-9-9-10-x
G6 Variant #21: -x-x-0-2-2-3
G6 Variant #22: -x-x-4-2-2-3
G6 Variant #23: -x-x-4-2-5-3
G6 Variant #24: -0-0-0-2-x-3
G6 Variant #25: -3-0-0-2-x-3
G6 Variant #26: -0-0-4-2-x-3
G6 Variant #27: -3-0-4-2-x-3
G6 Variant #28: -0-3-4-2-x-3
G6 Variant #29: -0-0-0-0-x-3
G6 Variant #30: -0-0-4-0-x-3
G6 Variant #31: -0-3-4-0-x-3
G6 Variant #32: -3-5-4-5-x-3
G6 Variant #33: -7-8-9-x-10-x
G6 Variant #34: -12-12-9-x-10-x
G6 Variant #35: -12-12-12-x-10-x
G6 Variant #36: -x-0-0-2-x-3
G6 Variant #37: -x-0-4-2-x-3
G6 Variant #38: -x-3-4-2-x-3
G6 Variant #39: -x-5-4-2-x-3
G6 Variant #40: -x-5-4-5-x-3
----------------------------------------------------------------------
G has 27 variations
Base chord:     -3-0-0-0-2-3
G Variant #0:   -3-3-0-0-2-3
G Variant #1:   -3-0-4-0-2-3
G Variant #2:   -3-3-4-5-5-3
G Variant #3:   -3-3-4-5-x-x
G Variant #4:   -7-8-7-5-x-x
G Variant #5:   -7-8-7-9-10-x
G Variant #6:   -10-12-12-12-10-x
G Variant #7:   -x-0-0-0-2-3
G Variant #8:   -x-3-0-0-2-3
G Variant #9:   -x-0-4-0-2-3
G Variant #10:  -x-3-4-0-2-3
G Variant #11:  -x-3-4-5-5-3
G Variant #12:  -x-8-7-9-10-x
G Variant #13:  -x-12-12-12-10-x
G Variant #14:  -x-x-0-0-2-3
G Variant #15:  -x-x-4-0-2-3
G Variant #16:  -x-x-4-5-5-3
G Variant #17:  -3-0-0-0-x-3
G Variant #18:  -3-0-4-0-x-3
G Variant #19:  -3-3-4-5-x-3
G Variant #20:  -7-8-7-x-10-x
G Variant #21:  -10-12-12-x-10-x
G Variant #22:  -x-0-0-0-x-3
G Variant #23:  -x-0-4-0-x-3
G Variant #24:  -x-3-4-0-x-3
G Variant #25:  -x-3-4-5-x-3
All done!
*/
