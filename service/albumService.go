package service

import (
	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/mappers"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/repository"
)

type AlbumService interface {
	GetAlbums() ([]generated.Album, error)
	CreateAlbum(string) (generated.Album, error)
}

type AlbumServiceImpl struct {
	repo repository.AlbumRepositoryInterface;
	albumMapper mappers.AlbumMapper;
}
func NewAlbumService(
	repo repository.AlbumRepositoryInterface,
	albumMapper mappers.AlbumMapper,
) AlbumService {
	return &AlbumServiceImpl{repo: repo, albumMapper: albumMapper}
}

func (a *AlbumServiceImpl) GetAlbums() ([]generated.Album, error) {
	albums, err := a.repo.FindAll()
	if err != nil {
		return nil, err  
	}
	return a.albumMapper.FromModelToGeneratedAlbums(albums), nil;
}

func (a *AlbumServiceImpl) CreateAlbum(albumnavn string) (generated.Album, error) {
	album := models.Album{Name: albumnavn}
	albumModel, err := a.repo.Create(&album)
	return  a.albumMapper.FromModelToGeneratedAlbum(albumModel), err 
}
