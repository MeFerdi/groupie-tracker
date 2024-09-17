package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadArtist(baseUrl, id string) (Artist, error) {
	baseUrl = fmt.Sprintf("%s%s", baseUrl, id)
	response, err := http.Get(baseUrl)
	if err != nil {
		return Artist{}, err
	}
	defer response.Body.Close()

	var artist Artist
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artist)
		if err != nil {
			return Artist{}, err
		}
		if artist.ID == 0{
			return Artist{}, fmt.Errorf("Artist not found")
		}
	} else {
		return Artist{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
