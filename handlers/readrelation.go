package api

import (
	"encoding/json"
	"net/http"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetRelationsData() (*Relation, error) {
	var relation Relation

	err := FetchData("https://groupietrackers.herokuapp.com/api/relation", &relation)
	if err != nil {
		return nil, err
	}

	return &relation, nil
}
