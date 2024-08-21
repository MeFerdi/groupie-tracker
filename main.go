package main

import (
	"net/http"

	artistApi "artistApi/handler"
)

func main() {
	http.HandleFunc("/artists", artistApi.ArtistHandler)
	http.ListenAndServe(":3000", nil)
}
