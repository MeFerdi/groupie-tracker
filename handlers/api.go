package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadApi() Location {
	url := "https://groupietrackers.herokuapp.com/api/locations/1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	locations := Location{}
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&locations)
		if err != nil {
			log.Fatal(err)
		}
	}
	return locations
}
