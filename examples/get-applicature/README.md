# Get applicature  

```go
// Create a new ultimateguitar
scraper := ultimateguitar.New()

// Fetch the applicatures for the chords for the instrument in the tuning specified
applicatures, err := scraper.GetApplicature("guitar", "E A D G B E", []ultimateguitar.ChordKey{
  ultimateguitar.ChordKeyA,
  ultimateguitar.ChordKeyG,
  ultimateguitar.ChordKeyG6,
})

if err != nil {
  // you should probably handle this lol.
}

// Iterate over the Applicature slice
for _, applicature := range applicatures {
  fmt.Println(applicature.Chord) // A, G, G6
}
```

See `./main.go` in this directory for an example. 