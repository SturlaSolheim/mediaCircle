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
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 204})
}
