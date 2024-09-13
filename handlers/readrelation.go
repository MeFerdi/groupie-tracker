package api

import (
	"encoding/json"
	"net/http"
)

func FetchRelations(id string) (Relation, error) {
	
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return Relation{}, err
	}
	defer res.Body.Close()

	
	var data Relation
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return Relation{}, err
	}

	return data, nil
}
