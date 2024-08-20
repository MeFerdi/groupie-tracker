package main

import (
	"net/http"

	api "groupie/handlers"
)

func main() {
	http.HandleFunc("/locations", api.LocationHandler)
	http.ListenAndServe(":3000", nil)
}
