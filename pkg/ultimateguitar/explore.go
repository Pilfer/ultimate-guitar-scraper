package ultimateguitar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ExploreOrder string

const (
	ExploreOrderTodaysMostPopular     ExploreOrder = "hitsdailygroup_desc" /* Today's most popular */
	ExploreOrderMostPopularOfAllTime  ExploreOrder = "hitstotal_desc"      /* Most popular of all time */
	ExploreOrderRecentlyAdded         ExploreOrder = "date_desc"           /* Recently added */
	ExploreOrderSongAZ                ExploreOrder = "songname_asc"        /* Song, A-Z */
	ExploreOrderSongZA                ExploreOrder = "songname_desc"       /* Song, Z-A */
	ExploreOrderArtistAZ              ExploreOrder = "artistname_asc"      /* Artist, A-Z */
	ExploreOrderArtistZA              ExploreOrder = "artistname_desc"     /* Artist, Z-A */
	ExploreOrderRating                ExploreOrder = "rating_desc"         /* Rating */
	ExploreOrderMostPopularOfAllTime2 ExploreOrder = "hits"                /* Most popular of all-time...alt */
	ExploreOrderTodaysMostPopular2    ExploreOrder = "hits_daily"          /* Most popular for today...alt */
)

// GetTabByID - Fetches the corresponding tab on UG
func (s *Scraper) Explore(params ExploreParameters) ([]ExploreResult, error) {
	tabResult := []ExploreResult{}
	v, err := query.Values(params)
	if err != nil {
		return tabResult, err
	}

	queryParams := v.Encode()
	urlString := fmt.Sprintf("%s%s?%s", ugAPIEndpoint, AppPaths.EXPLORE_TABS, queryParams)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return tabResult, err
	}

	s.ConfigureHeaders(req)

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

type ExploreParameters struct {
	Type       []TabType    `json:"type,omitempty" url:"type,omitempty"`             // Type of results to return
	Order      ExploreOrder `json:"order,omitempty" url:"order,omitempty"`           // Sort order
	Date       int32        `json:"date,omitempty" url:"date,omitempty"`             // I've only ever seen this set to 0. Not sure what it does.
	Page       int32        `json:"page,omitempty" url:"page,omitempty"`             // Page number to request
	Capo       []int32      `json:"capo,omitempty" url:"capo,omitempty"`             // 0 for No Capo, 1 for With Capo
	Decade     []int32      `json:"decade,omitempty" url:"decade,omitempty"`         // Decade the song was released. Ex: 1990, 2020, etc
	Tuning     []int32      `json:"tuning,omitempty" url:"tuning,omitempty"`         // Predefined tuning ids - starts at 1
	Genres     []int32      `json:"genres,omitempty" url:"genres,omitempty"`         // Predefined genre ids - starts at 1
	SubGenres  []int32      `json:"subgenres,omitempty" url:"subgenres,omitempty"`   // Predefined subgenre ids - starts at 1
	Difficulty []Difficulty `json:"difficulty,omitempty" url:"difficulty,omitempty"` // Difficulty of the song - 0 for Novice, 1 for Intermediate, 2 for Advance, 3 for None
}

type ExploreResult struct {
	ID                 int64         `json:"id"`
	SongID             int64         `json:"song_id"`
	SongName           string        `json:"song_name"`
	ArtistID           int64         `json:"artist_id"`
	ArtistName         string        `json:"artist_name"`
	Type               Type          `json:"type"`
	Part               string        `json:"part"`
	Version            int64         `json:"version"`
	Votes              int64         `json:"votes"`
	Rating             float64       `json:"rating"`
	Date               string        `json:"date"`
	Status             Status        `json:"status"`
	PresetID           int64         `json:"preset_id"`
	TabAccessType      TabAccessType `json:"tab_access_type"`
	TpVersion          int64         `json:"tp_version"`
	TonalityName       string        `json:"tonality_name"`
	VersionDescription *string       `json:"version_description"`
	Verified           int64         `json:"verified"`
	Recording          Recording     `json:"recording"`
}
