package main

import (
	"net/http"

	artistApi "artistApi/handler"
)

func main() {
	http.HandleFunc("/", artistApi.ArtistHandler)
	http.ListenAndServe(":8080", nil)
}
