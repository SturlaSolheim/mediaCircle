package handlers

import (
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/albums"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/anmeldelse"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/liste"
)

type CompositeHandler struct {
	albumHandler  albums.ServerInterface
	anmeldelseHandler  anmeldelse.ServerInterface
	listeHandler liste.ServerInterface
}

func NewCompositeHandler(
	albumHandler albums.ServerInterface,
	anmeldelseHandler anmeldelse.ServerInterface,
	listeHandler liste.ServerInterface,
) generated.ServerInterface {
	return &CompositeHandler{
		albumHandler:  albumHandler,
		anmeldelseHandler:  anmeldelseHandler,
		listeHandler: listeHandler,
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

// Liste
func (c *CompositeHandler) GetListe(w http.ResponseWriter, r *http.Request, id int) {
	c.listeHandler.GetListe(w, r, id)
}
func (c *CompositeHandler) GetListeTest(w http.ResponseWriter, r *http.Request, id int) {
	c.listeHandler.GetListeTest(w, r, id)
}
