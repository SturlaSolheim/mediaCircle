package handlers

import (
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/albums"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/anmeldelse"
)

type CompositeHandler struct {
	albumHandler  albums.ServerInterface
	anmeldelseHandler  anmeldelse.ServerInterface
}

func NewCompositeHandler(
	albumHandler albums.ServerInterface,
	anmeldelseHandler anmeldelse.ServerInterface,
) generated.ServerInterface {
	return &CompositeHandler{
		albumHandler:  albumHandler,
		anmeldelseHandler:  anmeldelseHandler,
	}
}

// Albums

func (c *CompositeHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.GetAlbums(w, r)
}

func (c *CompositeHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.CreateAlbum(w, r)
}

// Anmeldelse

func (c *CompositeHandler) GetAlleAnmeldelserForBruker(w http.ResponseWriter, r *http.Request, bruker string) {
	c.anmeldelseHandler.GetAlleAnmeldelserForBruker(w, r, bruker)
}
