package ultimateguitar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cheggaaa/pb/v3"
)

// GetAll - Fetches all the tabs/songs saved by user on UG
func (s *Scraper) GetAll() ([]TabResult, error) {
	tabListings := []TabListing{}
	tabResults := []TabResult{}
	if s.Token == "" {
		return tabResults, errors.New("GetAll(): requires token to fetch all tabs")
	}
	urlString := fmt.Sprintf("%s%s?token=%s", ugAPIEndpoint, AppPaths.MY_TABS, s.Token)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return tabResults, err
	}

	s.ConfigureHeaders(req)
	
	res, err := s.Client.Do(req)
	if err != nil {
		return tabResults, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&tabListings)
	if err != nil {
		return tabResults, err
	}


	bar := pb.Simple.Start(len(tabListings))
	for i := 0; i < len(tabListings); i++ {
		bar.Increment()
		tabResult, err := s.GetTabByID(int64(tabListings[i].Tab.ID))
		if err != nil {
			continue
		}
		tabResults = append(tabResults, tabResult)
	}
	bar.Finish()
	return tabResults, err
}

// GetTabByID - Fetches the corresponding tab on UG
func (s *Scraper) GetTabByID(tabID int64) (TabResult, error) {
	tabResult := TabResult{}

	urlString := fmt.Sprintf("%s%s?tab_id=%d&tab_access_type=private", ugAPIEndpoint, AppPaths.TAB_INFO, tabID)
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

// Fetch a tab by ID and return the raw string - this is a debug method that shouldn't ever be used.
func (s *Scraper) GetTabByIDRaw(tabID int64) (string, error) {

	urlString := fmt.Sprintf("%s%s?tab_id=%d&tab_access_type=private", ugAPIEndpoint, AppPaths.TAB_INFO, tabID)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return "", err
	}

	s.ConfigureHeaders(req)

	res, err := s.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	var b bytes.Buffer

	_, err = b.ReadFrom(res.Body)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
