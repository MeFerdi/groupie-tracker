package main

import (
	"net/http"

	api "groupie/handlers"
)

func main() {
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/locations/", api.LocationHandler)
	http.HandleFunc("/artists/", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	http.HandleFunc("/relation/", api.RelationHandler)
	http.HandleFunc("/dates/", api.DateHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.ListenAndServe(":8080", nil)
}
