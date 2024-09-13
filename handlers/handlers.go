package api

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

var errorTemplate *template.Template

func init() {
	var err error
	errorTemplate, err = template.ParseFiles("template/error.html")
	if err != nil {
		log.Fatal("Error parsing error template:", err)
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	err := errorTemplate.Execute(w, struct {
		Code    int
		Message string
	}{
		Code:    status,
		Message: message,
	})
	if err != nil {
		log.Printf("Error rendering error template: %v", err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	// Parse the homepage template
	temp, err := template.ParseFiles("template/home.html") // Ensure you have home.html in the template directory
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	// Execute the template and write to the response
	err = temp.Execute(w, nil) // No data is passed to the homepage template
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	temp1, err := template.ParseFiles("template/artists.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	result, err := ReadArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error fetching artists")
		return
	}

	err = temp1.Execute(w, result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/artist/") || len(strings.Split(r.URL.Path, "/")) != 3 {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/artist.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	result, err := ReadArtist(id)
	if err != nil {
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "cannot unmarshal") {
			renderError(w, http.StatusNotFound, "Artist not found")
		} else {
			renderError(w, http.StatusInternalServerError, "Error fetching artist: "+err.Error())
		}
		return
	}

	if result.ID == 0 {
		renderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	err = temp1.Execute(w, result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/locations.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	Result, err := ReadLocation(id)
	if err != nil {
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "cannot unmarshal") {
			renderError(w, http.StatusNotFound, "Location not found")
		} else {
			renderError(w, http.StatusInternalServerError, "Error fetching location: "+err.Error())
		}
		return
	}

	if Result.ID == 0 {
		renderError(w, http.StatusNotFound, "Location not found")
		return
	}

	err = temp1.Execute(w, Result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/dates/") || len(strings.Split(r.URL.Path, "/")) != 3 {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/dates.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	Result, err := ReadDate(id)
	if err != nil {
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "cannot unmarshal") {
			renderError(w, http.StatusNotFound, "Date entry not found")
		} else {
			renderError(w, http.StatusInternalServerError, "Failed to fetch date entry: "+err.Error())
		}
		return
	}

	if Result.ID == 0 {
		renderError(w, http.StatusNotFound, "Date entry not found")
		return
	}

	err = temp1.Execute(w, Result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/relation/") || len(strings.Split(r.URL.Path, "/")) != 3 {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "ID not found")
		return
	}
	id := id1[len(id1)-1]

	relationTemplate, err := template.ParseFiles("template/relation.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template: "+err.Error())
		return
	}

	relations, err := FetchRelations(id)
	if err != nil {
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "cannot unmarshal") {
			renderError(w, http.StatusNotFound, "Relation not found")
		} else {
			renderError(w, http.StatusInternalServerError, "Failed to fetch relations data: "+err.Error())
		}
		return
	}

	if relations.ID == 0 {
		renderError(w, http.StatusNotFound, "Relation not found")
		return
	}

	err = relationTemplate.Execute(w, relations)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Failed to render template: "+err.Error())
	}
}
