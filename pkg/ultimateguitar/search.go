package ultimateguitar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

/*
title
page
type
official
tuning=E A D G B E
display_songs
artist_name
filter
*/
type SearchParams struct {
	Title  string    `json:"title,omitempty" url:"title,omitempty"`
	Type   []TabType `json:"type,omitempty" url:"type,omitempty"`
	Page   int32     `json:"page,omitempty" url:"page,omitempty"`
	Tuning string    `json:"tuning,omitempty" url:"tuning,omitempty"`
}

// Search for tabs on UG
func (s *Scraper) Search(params SearchParams) (SearchResult, error) {
	searchResult := SearchResult{}
	v, err := query.Values(params)
	if err != nil {
		return searchResult, err
	}

	queryParams := v.Encode()
	urlString := fmt.Sprintf("%s%s?%s", ugAPIEndpoint, AppPaths.SEARCH, queryParams)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return searchResult, err
	}

	s.ConfigureHeaders(req)

	res, err := s.Client.Do(req)
	if err != nil {
		return searchResult, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&searchResult)
	if err != nil {
		return searchResult, err
	}

	return searchResult, nil
}
