package ultimateguitar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetApplicatureParams struct {
	Tuning     string `json:"tuning,omitempty" url:"tuning,omitempty"`         // E A D G B E, etc.
	Instrument string `json:"instrument,omitempty" url:"instrument,omitempty"` // guitar, ukulele, etc
}

func (s *Scraper) GetApplicature(instrument string, tuning string, chords []ChordKey) ([]Applicature, error) {
	results := []Applicature{}
	v, _ := query.Values(GetApplicatureParams{Tuning: tuning, Instrument: instrument})
	for _, chord := range chords {
		v.Add(fmt.Sprintf("chords[%s]", chord), string(chord))
	}
	queryParams := v.Encode()
	// params := "tuning=E A D G B E&instrument=guitar&chords[Ab]=Ab"
	urlString := fmt.Sprintf("%s%s?%s", ugAPIEndpoint, AppPaths.APPLICATURE, queryParams)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return results, err
	}

	s.ConfigureHeaders(req)

	res, err := s.Client.Do(req)
	if err != nil {
		return results, err
	}

	defer res.Body.Close()

	// var responseBytes bytes.Buffer
	// _, err = responseBytes.ReadFrom(res.Body)
	// if err != nil {
	// 	return results, err
	// }
	// fmt.Println(responseBytes.String())

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}
