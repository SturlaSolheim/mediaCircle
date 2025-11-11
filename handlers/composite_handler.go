package handlers

import (
	"net/http"
	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/albums"
)

type CompositeHandler struct {
	albumHandler  albums.ServerInterface
}

func NewCompositeHandler(
	albumHandler albums.ServerInterface,
) generated.ServerInterface {
	return &CompositeHandler{
		albumHandler:  albumHandler,
	}
}

// Albums

func (c *CompositeHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.GetAlbums(w, r)
}

func (c *CompositeHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.CreateAlbum(w, r)
}

