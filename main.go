package main

import (
	"github.com/gorilla/mux"
	"go-albums/apis"
	handler "go-albums/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	artist := &apis.HttpArtist{}
	r.HandleFunc("/albums", handler.GetAlbums(artist))

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		return
	}



}