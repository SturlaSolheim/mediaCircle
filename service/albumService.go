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
	albums, err := a.repo.FindAll()
	if err != nil {
		return nil, err  
	}
	return albums, nil
}

func (a *AlbumService) CreateAlbum(album models.Album) error {
	return a.repo.Save(&album)
}
