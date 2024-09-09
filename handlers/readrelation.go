package api

import (
	"encoding/json"
	"net/http"
)

func FetchRelations(id string) (Relation, error) {
	// Make a GET request to the API endpoint for relation data
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return Relation{}, err
	}
	defer res.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a RelationData struct
	var data Relation
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return Relation{}, err
	}

	return data, nil
}
