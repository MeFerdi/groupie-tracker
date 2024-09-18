package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchRelations(baseURL, id string) (Relation, error) {
	// Make a GET request to the API endpoint for relation data
	res, err := http.Get(baseURL + id)
	if err != nil {
		return Relation{}, err
	}
	defer res.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a RelationData struct
	var data Relation
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return Relation{}, err
	}
	if data.ID == 0 {
		return Relation{}, fmt.Errorf("relation not found for ID: %s", id)
	}

	return data, nil
}
