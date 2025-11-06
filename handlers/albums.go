package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/go-chi/chi/v5"
)

var albums = []models.Album{
	{ID: 1, Name: "Abbey road"},
	{ID: 2, Name: "Schmillson"},
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	for _, album := range albums {
		if album.ID == parseID(id) {
			json.NewEncoder(w).Encode(album)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Album not found", Status: 204})
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{Message: "Invalid JSON", Status: 400})
		return
	}

	album.ID = len(albums) + 1
	albums = append(albums, album)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(albums)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	for i, album := range albums {
		if album.ID == parseID(id) {
			var updatedAlbum models.Album
			if err := json.NewDecoder(r.Body).Decode(&updatedAlbum); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.Response{Message: "Invalid JSON", Status: 400})
				return
			}
			updatedAlbum.ID = album.ID
			albums[i] = updatedAlbum
			json.NewEncoder(w).Encode(updatedAlbum)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Album not found", Status: 204})
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	for i, album := range albums {
		if album.ID == parseID(id) {
			albums = append(albums[:i], albums[i+1:]...)
			json.NewEncoder(w).Encode(models.Response{Message: "Album deleted", Status: 200})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Album not found", Status: 204})
}

func parseID(id string) int {
	if id == "1" {
		return 1
	} else if id == "2" {
		return 2
	}
	return 0
}
