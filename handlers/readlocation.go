package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadLocation(baseURL,id string) (Location, error) {
	baseURL = fmt.Sprintf("%s%s", baseURL, id)
	response, err := http.Get(baseURL)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	var artist Location
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artist)
		if err != nil {
			return Location{}, err
		}
	} else {
		return Location{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
