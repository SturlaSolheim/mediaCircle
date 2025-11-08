package routes

import (
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/handlers"
	"github.com/SturlaSolheim/mediaCircleBackend/repository"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
	"github.com/go-chi/chi/v5"
)

type Container struct {
	AlbumHandler *handlers.AlbumHandler
}

func NewContainer() *Container {
	// Repositories
	albumRepo := repository.NewAlbumRepository(database.DB)

	// Services
	albumService := service.NewAlbumService(albumRepo) 

	// Handlers
	albumHandler := handlers.NewAlbumHandler(*albumService)
	
	return &Container{
		AlbumHandler: albumHandler,
	}
}

func (c *Container) SetupRoutes(r chi.Router) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheck)
		r.Route("/albums", func(r chi.Router) {
			r.Get("/", c.AlbumHandler.GetAlbums)
			r.Post("/", c.AlbumHandler.CreateAlbum)
			r.Get("/{id}", c.AlbumHandler.GetAlbum)
			r.Put("/{id}", c.AlbumHandler.UpdateAlbum)
			r.Delete("/{id}", c.AlbumHandler.DeleteAlbum)
		})
	})
}
