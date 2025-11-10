package handlers

import (
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/albums"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/health"
)

type CompositeHandler struct {
	albumHandler  albums.ServerInterface
	healthHandler health.ServerInterface
}

func NewCompositeHandler(
	albumHandler albums.ServerInterface,
	healthHandler health.ServerInterface,
) generated.ServerInterface {
	return &CompositeHandler{
		albumHandler:  albumHandler,
		healthHandler: healthHandler,
	}
}

// Albums

func (c *CompositeHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.GetAlbums(w, r)
}

func (c *CompositeHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	c.albumHandler.CreateAlbum(w, r)
}

// Health

func (c *CompositeHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	c.healthHandler.HealthCheck(w, r)
}
