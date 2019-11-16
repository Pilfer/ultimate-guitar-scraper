package ultimateguitar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTabByID - Fetches the corresponding tab on UG
func (s *Scraper) GetTabByID(tabID int64) (TabResult, error) {
	tabResult := TabResult{}

	urlString := fmt.Sprintf("%s%s?tab_id=%d&tab_access_type=private", ugAPIEndpoint, AppPaths.TAB_INFO, tabID)
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

// TabResult struct - this is what we get when we fetch a tab by id.
type TabResult struct {
	ID                 int         `json:"id"`
	SongName           string      `json:"song_name"`
	ArtistName         string      `json:"artist_name"`
	Type               string      `json:"type"`
	Part               string      `json:"part"`
	Version            int         `json:"version"`
	Votes              int         `json:"votes"`
	Rating             float64     `json:"rating"`
	Date               string      `json:"date"`
	Status             string      `json:"status"`
	PresetID           int         `json:"preset_id"`
	TabAccessType      string      `json:"tab_access_type"`
	TpVersion          int         `json:"tp_version"`
	TonalityName       string      `json:"tonality_name"`
	VersionDescription interface{} `json:"version_description"`
	Verified           int         `json:"verified"`
	Recording          struct {
		IsAcoustic       int           `json:"is_acoustic"`
		TonalityName     string        `json:"tonality_name"`
		Performance      interface{}   `json:"performance"`
		RecordingArtists []interface{} `json:"recording_artists"`
	} `json:"recording"`
	Versions []struct {
		ID                 int     `json:"id"`
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
	} `json:"versions"`
	Recommended []struct {
		ID                 int         `json:"id"`
		SongName           string      `json:"song_name"`
		ArtistName         string      `json:"artist_name"`
		Type               string      `json:"type"`
		Part               string      `json:"part"`
		Version            int         `json:"version"`
		Votes              int         `json:"votes"`
		Rating             float64     `json:"rating"`
		Date               string      `json:"date"`
		Status             string      `json:"status"`
		PresetID           int         `json:"preset_id"`
		TabAccessType      string      `json:"tab_access_type"`
		TpVersion          int         `json:"tp_version"`
		TonalityName       string      `json:"tonality_name"`
		VersionDescription interface{} `json:"version_description"`
		Verified           int         `json:"verified"`
		Recording          struct {
			IsAcoustic       int           `json:"is_acoustic"`
			TonalityName     string        `json:"tonality_name"`
			Performance      interface{}   `json:"performance"`
			RecordingArtists []interface{} `json:"recording_artists"`
		} `json:"recording"`
	} `json:"recommended"`
	UserRating  int           `json:"userRating"`
	Difficulty  string        `json:"difficulty"`
	Tuning      string        `json:"tuning"`
	Capo        int           `json:"capo"`
	URLWeb      string        `json:"urlWeb"`
	Strumming   []interface{} `json:"strumming"`
	VideosCount int           `json:"videosCount"`
	ProBrother  interface{}   `json:"pro_brother"`
	Contributor struct {
		UserID   int    `json:"user_id"`
		Username string `json:"username"`
	} `json:"contributor"`
	Applicature []struct {
		Chord      string `json:"chord"`
		Variations []struct {
			ID        string        `json:"id"`
			ListCapos []interface{} `json:"listCapos"`
			NoteIndex int           `json:"noteIndex"`
			Notes     []int         `json:"notes"`
			Frets     []int         `json:"frets"`
			Fingers   []int         `json:"fingers"`
			Fret      int           `json:"fret"`
		} `json:"variations"`
	} `json:"applicature"`
	Content string `json:"content"`
}
