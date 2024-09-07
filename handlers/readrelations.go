package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchData(url string, target interface{}) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error decoding data from %s: %w", url, err)
	}

	return nil
}

func getRelationData(artistID string) (*PageData, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", artistID)
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", artistID)

	var artist Artist
	var relation Relation

	// Fetch artist data
	if err := fetchData(artistURL, &artist); err != nil {
		return nil, err
	}

	// Fetch relation data
	if err := fetchData(relationURL, &relation); err != nil {
		return nil, err
	}

	// Debugging: Print the fetched data
	log.Printf("Fetched artist data: %+v", artist)
	log.Printf("Fetched relation data: %+v", relation)

	return &PageData{Artist: artist, Relation: relation}, nil
}
