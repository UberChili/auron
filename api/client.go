package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const SINGLE_TERM_SEARCH_URL = "https://aur.archlinux.org/rpc/v5/search/"

type SearchResponse struct {
	Resultcount int        `json:"resultcount"`
	Results     []PckgInfo `json:"results"`
	Type        string     `json:"type"`
}

type PckgInfo struct {
	Description string `json:"Description"`
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Url         string `json:"URL"`
	Version     string `json:"Version"`
}

// searches for a package
func SearchPackage(packageName string) (*SearchResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", SINGLE_TERM_SEARCH_URL+packageName, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("by", "name")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch results. Status code %d\n", resp.StatusCode)
	}

	var results SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return &results, nil
}
