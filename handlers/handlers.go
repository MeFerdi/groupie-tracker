package api

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

var relationTemplate *template.Template

// Initialize the template
func init() {
	var err error
	relationTemplate, err = template.ParseFiles("template/relation.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}

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
		log.Println("Error executing template:", err)
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

	if result.ID == 0 { 
        http.NotFound(w, r) 
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

	Result, err := ReadDate(id)
	if err != nil {
        http.Error(w, "Failed to fetch date entry: "+err.Error(), http.StatusInternalServerError)
        return
    }
	if Result.ID == 0 { 
        http.NotFound(w, r) 
        return
    }
	err = temp1.Execute(w, Result)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the artist ID from the URL
	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}
	id := id1[len(id1)-1]

	// Fetch the relations data
	relations, err := FetchRelations(id)
	if err != nil {
		log.Printf("Error fetching relations data: %v", err)
		http.Error(w, "Failed to fetch relations data", http.StatusInternalServerError)
		return
	}

	// Check if relationTemplate is properly initialized
	if relationTemplate == nil {
		log.Println("Error: relationTemplate is nil")
		http.Error(w, "Template is not initialized", http.StatusInternalServerError)
		return
	}

	// Render the template with relations data
	err = relationTemplate.Execute(w, relations)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
