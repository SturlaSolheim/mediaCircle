package repository

import (
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/gorm"
)

type AlbumRepositoryInterface interface {
	Save(album *models.Album) error
	FindAll() ([]models.Album, error)
}

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) AlbumRepositoryInterface {
	return &albumRepository{db: db}
}

func (r *albumRepository) Save(album *models.Album) error {
	return r.db.Save(album).Error
}

func (r *albumRepository) FindAll() ([]models.Album, error) {
	var albums []models.Album
	err := r.db.Find(&albums).Error
	return albums, err
}
