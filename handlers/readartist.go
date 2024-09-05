package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadArtist(id string) (Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists/"
	url = fmt.Sprintf("%s%s", url, id)
	response, err := http.Get(url)
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
	} else {
		return Artist{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
