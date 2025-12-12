package service

import (
	"testing"

	"github.com/SturlaSolheim/mediaCircleBackend/mappers"
	mocks "github.com/SturlaSolheim/mediaCircleBackend/mocks"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/stretchr/testify/assert"
)

func TestAlbumService_GetAlbums(t *testing.T) {
	setup := func() (*mocks.MockAlbumRepositoryInterface, AlbumService) {
		mockRepo := mocks.NewMockAlbumRepositoryInterface(t)
		albumMapper := mappers.NewAlbumMapper()
		albumService := NewAlbumService(mockRepo, *albumMapper)
		return mockRepo, albumService
	}

	t.Run("Returnerer album riktig", func(t *testing.T) {
		mockRepo, albumService := setup()

		expectedAlbums := []models.Album{
			{ID: uint(1), Name: "Test Album 1"},
			{ID: uint(2), Name: "Test Album 2"},
		}

		mockRepo.EXPECT().FindAll().Return(expectedAlbums, nil)

		result, err := albumService.GetAlbums()

		assert.NoError(t, err)
		assert.Equal(t, len(expectedAlbums), len(result))
		assert.Equal(t, expectedAlbums[0].Name, result[0].Name)
		assert.Equal(t, expectedAlbums[0].ID, uint(result[0].Id))

		assert.Equal(t, expectedAlbums[1].Name, result[1].Name)
		assert.Equal(t, expectedAlbums[1].ID, uint(result[1].Id))
	})
}

func TestAlbumService_CreateAlbum(t *testing.T){
	setup := func() (*mocks.MockAlbumRepositoryInterface, AlbumService) {
		mockRepo := mocks.NewMockAlbumRepositoryInterface(t)
		albumMapper := mappers.NewAlbumMapper()
		albumService := NewAlbumService(mockRepo, *albumMapper)
		return mockRepo, albumService
	}
	t.Run("Fungerer riktig", func(t *testing.T) {
		mockRepo, albumService := setup()

		expectedAlbum := models.Album{ID: uint(1), Name: "Abbey Road"}

		mockRepo.EXPECT().Create(&models.Album{Name: "Abbey Road"}).Return(expectedAlbum, nil)

		result, err := albumService.CreateAlbum("Abbey Road")
		assert.NoError(t, err)
		assert.Equal(t, expectedAlbum.Name, result.Name)
		assert.Equal(t, expectedAlbum.ID, uint(result.Id))
	})
}
