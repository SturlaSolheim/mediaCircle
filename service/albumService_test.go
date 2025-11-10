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
			{ID: 1, Name: "Test Album 1"},
			{ID: 2, Name: "Test Album 2"},
		}

		mockRepo.EXPECT().FindAll().Return(expectedAlbums, nil)

		result, err := albumService.GetAlbums()

		assert.NoError(t, err)
		assert.Equal(t, len(expectedAlbums), len(result))
		assert.Equal(t, expectedAlbums[0].Name, result[0].Name)
		assert.Equal(t, expectedAlbums[0].ID, result[0].Id)

		assert.Equal(t, expectedAlbums[1].Name, result[1].Name)
		assert.Equal(t, expectedAlbums[1].ID, result[1].Id)
	})

}
