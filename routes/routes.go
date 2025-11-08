package routes

import (
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/handlers"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	albumRepo := database.NewAlbumRepository()
	albumService := service.NewAlbumService(*albumRepo)
	albumHandler := handlers.NewAlbumHandler(*albumService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheck)
		r.Route("/albums", func(r chi.Router) {
			r.Get("/", albumHandler.GetAlbums)
			r.Post("/", albumHandler.CreateAlbum)
			r.Get("/{id}", handlers.GetAlbum)
			r.Put("/{id}", handlers.UpdateAlbum)
			r.Delete("/{id}", handlers.DeleteAlbum)
		})
	})
}
