package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type Artist struct {
// 	Name       string   `json:"name"`
// 	Image      string   `json:"image"`
// 	Year       int      `json:"year"`
// 	FirstAlbum string   `json:"firstAlbum"`
// 	Members    []string `json:"members"`
// }

type Relation struct {
	ArtistName string   `json:"artist_name"`
	Locations  []string `json:"locations"`
	Dates      []string `json:"dates"`
}

type PageData struct {
	Artist   Artist
	Relation Relation
}

func getRelationData(artistID string) (*PageData, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", artistID)
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", artistID)

	var artist Artist
	var relation Relation

	// Fetch artist info
	artistResp, err := http.Get(artistURL)
	if err != nil {
		return nil, err
	}
	defer artistResp.Body.Close()

	if err := json.NewDecoder(artistResp.Body).Decode(&artist); err != nil {
		return nil, err
	}

	// Fetch relation info
	relationResp, err := http.Get(relationURL)
	if err != nil {
		return nil, err
	}
	defer relationResp.Body.Close()

	if err := json.NewDecoder(relationResp.Body).Decode(&relation); err != nil {
		return nil, err
	}

	return &PageData{Artist: artist, Relation: relation}, nil
}
