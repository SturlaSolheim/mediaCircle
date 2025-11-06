package database

import (
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/gorm"
)

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{db: DB}
}

func (r *AlbumRepository) Save(album *models.Album) error {
	return r.db.Save(album).Error
}

func (r *AlbumRepository) FindByID(id uint) (*models.Album, error) {
	var album models.Album
	err := r.db.First(&album, id).Error
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (r *AlbumRepository) FindAll() ([]models.Album, error) {
	var albums []models.Album
	err := r.db.Find(&albums).Error
	return albums, err
}

func (r *AlbumRepository) FindByName(name string) ([]models.Album, error) {
	var albums []models.Album
	err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&albums).Error
	return albums, err
}

func (r *AlbumRepository) Delete(id uint) error {
	return r.db.Delete(&models.Album{}, id).Error
}

func (r *AlbumRepository) DeleteByEntity(album *models.Album) error {
	return r.db.Delete(album).Error
}

func (r *AlbumRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Album{}).Count(&count).Error
	return count, err
}

