package main

import (
	"net/http"

	api "groupie/handlers"
)

func main() {
	http.HandleFunc("/locations/", api.LocationHandler)
	http.HandleFunc("/", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	http.HandleFunc("/relation/", api.RelationHandler)
	http.HandleFunc("/dates/", api.DateHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.ListenAndServe(":3000", nil)
}
