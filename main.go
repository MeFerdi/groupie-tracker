package main

import (
	"net/http"

	api "groupie/handlers"
)

func main() {
	http.HandleFunc("/locations/", api.LocationHandler)
	http.HandleFunc("/", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	http.HandleFunc("/relation/", api.RelationsHandler)
	http.HandleFunc("/dates/", api.DateHandler)

	http.ListenAndServe(":3000", nil)
}
