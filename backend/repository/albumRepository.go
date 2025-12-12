package repository

import (
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/gorm"
)

type AlbumRepositoryInterface interface {
	Create(album *models.Album) (models.Album, error)
	FindAll() ([]models.Album, error)
}

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) AlbumRepositoryInterface {
	return &albumRepository{db: db}
}

func (r *albumRepository) Create(album *models.Album) (models.Album, error) {
	result := r.db.Create(&album)
	return *album, result.Error
}

func (r *albumRepository) FindAll() ([]models.Album, error) {
	var albums []models.Album
	err := r.db.Find(&albums).Error
	return albums, err
}
