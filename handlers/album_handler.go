package handlers

import (
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
	albums, err := h.albumService.GetAlbums()
	if err != nil {
		WriteInternalServerError(w, "Internal server error")
		return
	}
	
	WriteJSON(w, http.StatusOK, albums)
}

func (h *AlbumHandlerImpl) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var request generated.CreateAlbumRequest
	if err := DecodeJSON(r, &request); err != nil {
		WriteBadRequest(w, "Invalid JSON")
		return
	}
	
	album, err := h.albumService.CreateAlbum(request.Name)
	if err != nil {
		WriteInternalServerError(w, "Internal server error")
		return
	}
	
	WriteJSON(w, http.StatusCreated, album)
}
