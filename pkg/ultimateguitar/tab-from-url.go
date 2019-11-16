package ultimateguitar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Scraper) TabByURL(urlParam string) (TabFromURLResult, error) {
	var tabResult TabFromURLResult

	urlString := fmt.Sprintf("%s/tab/url?url=%s", ugAPIEndpoint, urlParam)
	req, _ := http.NewRequest("GET", urlString, nil)

	// Do some minor header manipulation so we retain the case
	for key := range ugHeaders {
		req.Header[key] = []string{ugHeaders[key]}
	}
	req.Header["X-UG-CLIENT-ID"] = []string{s.DeviceID}
	req.Header["X-UG-API-KEY"] = []string{s.generateAPIKey()}

	// This header isn't sent in the app, so we remove it.
	req.Header.Del("Accept-Encoding")

	res, err := s.Client.Do(req)
	if err != nil {
		return tabResult, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&tabResult)
	if err != nil {
		return tabResult, err
	}

	return tabResult, nil
}

type TabFromURLResult struct {
	ID                 int     `json:"id"`
	SongID             int     `json:"song_id"`
	SongName           string  `json:"song_name"`
	ArtistName         string  `json:"artist_name"`
	Type               string  `json:"type"`
	Part               string  `json:"part"`
	Version            int     `json:"version"`
	Votes              int     `json:"votes"`
	Rating             float64 `json:"rating"`
	Date               string  `json:"date"`
	Status             string  `json:"status"`
	PresetID           int     `json:"preset_id"`
	TabAccessType      string  `json:"tab_access_type"`
	TpVersion          int     `json:"tp_version"`
	TonalityName       string  `json:"tonality_name"`
	VersionDescription string  `json:"version_description"`
	Verified           int     `json:"verified"`
	Recording          struct {
		IsAcoustic       int           `json:"is_acoustic"`
		TonalityName     string        `json:"tonality_name"`
		Performance      interface{}   `json:"performance"`
		RecordingArtists []interface{} `json:"recording_artists"`
	} `json:"recording"`
}
