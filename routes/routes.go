package routes

import (
	"github.com/SturlaSolheim/mediaCircleBackend/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheck)
		r.Route("/albums", func(r chi.Router) {
			r.Get("/", handlers.GetAlbums)
			r.Post("/", handlers.CreateAlbum)
			r.Get("/{id}", handlers.GetAlbum)
			r.Put("/{id}", handlers.UpdateAlbum)
			r.Delete("/{id}", handlers.DeleteAlbum)
		})
	})
}
