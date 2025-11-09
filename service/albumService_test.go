package service

import (
	"testing"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
	mocks "github.com/SturlaSolheim/mediaCircleBackend/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAlbumService_GetAlbums(t *testing.T) {
	setup := func() (*mocks.MockAlbumRepositoryInterface, *AlbumService) {
		mockRepo := mocks.NewMockAlbumRepositoryInterface(t)
		albumService := NewAlbumService(mockRepo)
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
		assert.Equal(t, expectedAlbums, result)
	})

}
