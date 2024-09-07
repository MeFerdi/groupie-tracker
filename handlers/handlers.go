package api

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	temp1, err := template.ParseFiles("template/artists.html")
	if err != nil {
		log.Println("Error loading template:", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	result, err := ReadArtists()
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	// Pass the result to the template
	err = temp1.Execute(w, result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the artist ID from the URL
	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		http.Error(w, "Artist ID not found", http.StatusBadRequest)
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/artist.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Fetch the artist data
	result, err := ReadArtist(id)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	// Pass the result to the template
	err = temp1.Execute(w, result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}
	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		http.Error(w, "Artist ID not found", http.StatusBadRequest)
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/locations.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	Result, _ := ReadLocation(id)

	// Pass the result to the template
	err = temp1.Execute(w, Result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}
	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		http.Error(w, "Artist ID not found", http.StatusBadRequest)
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/dates.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	Result, _ := ReadDate(id)

	err = temp1.Execute(w, Result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

}
func RelationsHandler(w http.ResponseWriter, r *http.Request) {

	idSegments := strings.Split(r.URL.Path, "/")
	if len(idSegments) < 3 {
		http.Error(w, "Artist ID not found", http.StatusBadRequest)
		return
	}
	artistID := idSegments[len(idSegments)-1]

	pageData, err := getRelationData(artistID)
	if err != nil {
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("template/relation.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
