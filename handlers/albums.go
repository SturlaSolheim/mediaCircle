package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	// "github.com/go-chi/chi/v5"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	repo := database.NewAlbumRepository()
	albums, err := repo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	json.NewEncoder(w).Encode(albums)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")


	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	
	if album.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Album name is required",
			Status:  400,
		})
		return
	}
	
	repo := database.NewAlbumRepository()
	err = repo.Save(&album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Message: "En feil skjedde",
			Status:  500,
		})
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}
