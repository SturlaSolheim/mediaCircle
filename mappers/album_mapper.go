
package mappers

import (
	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
)

type AlbumMapper struct {}

func NewAlbumMapper() *AlbumMapper {
	return &AlbumMapper{}
}

func (a *AlbumMapper) FromModelToGeneratedAlbum(album models.Album) generated.Album {
	return generated.Album{
		Id:        int64(album.ID),
		Name:      album.Name,
		CreatedAt: album.CreatedAt,
		UpdatedAt: album.UpdatedAt,
	}
}

func (a *AlbumMapper)FromModelToGeneratedAlbums(albums []models.Album) []generated.Album {
	result := make([]generated.Album, len(albums))
	for i, album := range albums {
		result[i] = a.FromModelToGeneratedAlbum(album)
	}
	return result
}

func (a *AlbumMapper)ErrorResponse(message string, status int) generated.Response {
	return generated.Response{
		Message: message,
		Status:  status,
	}
}
