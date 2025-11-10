package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/albums"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
)

type AlbumHandlerImpl struct {
	albumService service.AlbumService
}

func NewOpenAPIAlbumHandler(
	albumService service.AlbumService,
) albums.ServerInterface {
	return &AlbumHandlerImpl{
		albumService: albumService,
	}
}

func (h *AlbumHandlerImpl) GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	albums, err := h.albumService.GetAlbums()
	if err != nil {
		h.returnInternalServerError(w, "Internal server error")
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func (h *AlbumHandlerImpl) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var req generated.CreateAlbumRequest;
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.returnBadRequest(w, "Invalid JSON")
		return
	}
	
	album, err := h.albumService.CreateAlbum(req.Name)
	if   err != nil {
		h.returnInternalServerError(w, "Internal server error")
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}


func (h *AlbumHandlerImpl) returnBadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	response := generated.Response{Message: message, Status: 400}
	json.NewEncoder(w).Encode(response)
}

func (h *AlbumHandlerImpl) returnInternalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	response := generated.Response{Message: message, Status: 500}
	json.NewEncoder(w).Encode(response)
}
