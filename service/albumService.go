package service

import (
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
)
type AlbumService struct {
	repo *database.AlbumRepository
}
func NewAlbumService() *AlbumService {
    return &AlbumService{
        repo: database.NewAlbumRepository(),
    }
}

func (a *AlbumService) GetAlbums() ([]models.Album, error) {
	repo := database.NewAlbumRepository()
	albums, err := repo.FindAll()
	if err != nil {
		return nil, err  
	}
	return albums, nil
}
