package handler

import (
	"encoding/json"
	"go-albums/apis"
	"go-albums/models"
	"net/http"
)

var URL = "https://jsonplaceholder.typicode.com/albums"

func GetAlbums(artist apis.HttpInterface) http.HandlerFunc  {
	return func (w http.ResponseWriter, r *http.Request) {
		var albums []models.Album

		body, err := artist.Get(URL)
		returnError(w, err)

		err = json.Unmarshal(body, &albums)
		returnError(w, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(albums)
		if err != nil {
			return
		}
	}
}

func returnError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}