package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
	// "github.com/go-chi/chi/v5"
)

type AlbumHandler struct {
	albumService *service.AlbumService
}

func NewAlbumHandler(albumService service.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		albumService: &albumService,
	}
}

func (h *AlbumHandler)GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	albums, err := h.albumService.GetAlbums()
		if err != nil {
			returnInternalServerError(w)
			return
		}
	json.NewEncoder(w).Encode(albums)
}

func (h *AlbumHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func (h *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		returnBadRequest(w)
		return
	}
	
	if album.Name == "" {
		returnBadRequest(w)
		return
	}
	
	err = h.albumService.CreateAlbum(album)
	if err != nil {
		returnInternalServerError(w)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

func (h *AlbumHandler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func (h *AlbumHandler)DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}


func returnBadRequest(w http.ResponseWriter) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Invalid request",
			Status:  400,
		})
}

func returnInternalServerError(w http.ResponseWriter){
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Message: "En feil skjedde",
			Status:  500,
		})
}

func returnNotImplemented(w http.ResponseWriter){
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(models.Response{Message: "Ikke implementert enda", Status: 501})
}
