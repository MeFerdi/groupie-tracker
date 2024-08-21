package artistApi

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Artist struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	FirstAlbum string   `json:"first_album"`
	Members    []string `json:"members"`
}

type Artists []Artist

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
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

	err = temp1.Execute(w, Result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func ReadApi() Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var artists Artists
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artists)
		if err != nil {
			log.Fatal(err)
		}
	}
	return artists
}
