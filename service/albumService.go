package service

import (
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
)

func GetAlbums() ([]models.Album, error) {
	repo := database.NewAlbumRepository()
	albums, err := repo.FindAll()
	if err != nil {
		return nil, err  
	}
	return albums, nil
}
