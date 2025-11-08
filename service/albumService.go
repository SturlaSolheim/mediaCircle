package service

import (
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/repository"
)
type AlbumService struct {
	repo repository.AlbumRepositoryInterface
}
func NewAlbumService(repo repository.AlbumRepositoryInterface) *AlbumService {
	return &AlbumService{repo: repo}
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
