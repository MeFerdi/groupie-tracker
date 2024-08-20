package api

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Location struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return 
	}

	temp1, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	Result := ReadApi()

	// Pass the result to the template
	err = temp1.Execute(w, Result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

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
