package routes

import (
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/handlers"
	"github.com/SturlaSolheim/mediaCircleBackend/mappers"
	"github.com/SturlaSolheim/mediaCircleBackend/repository"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
	"github.com/go-chi/chi/v5"
)

type OpenAPIContainer struct {
	handler generated.ServerInterface
}

func NewOpenAPIContainer() *OpenAPIContainer {
	// Repositories
	albumRepository := repository.NewAlbumRepository(database.DB)

	// Mappers
	albumMapper := mappers.NewAlbumMapper()

	// Services
	albumService := service.NewAlbumService(albumRepository, *albumMapper)

	// Handlers
	albumHandler := handlers.NewOpenAPIAlbumHandler(albumService)
	healthHandler := handlers.NewOpenAPIHealthHandler()

	// Composite
	compositeHandler := handlers.NewCompositeHandler(
		albumHandler,
		healthHandler,
	)

	return &OpenAPIContainer{
		handler: compositeHandler,
	}
}

func (c *OpenAPIContainer) SetupOpenAPIRoutes(r chi.Router) {
	generated.HandlerFromMuxWithBaseURL(c.handler, r, "/api/v1")
}
