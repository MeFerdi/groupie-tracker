package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []Artist
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artists)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artists, nil
}
