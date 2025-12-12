package service

import (
	"testing"

	"github.com/SturlaSolheim/mediaCircleBackend/mappers"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/repository"
	"github.com/SturlaSolheim/mediaCircleBackend/test"
	"github.com/stretchr/testify/assert"
)

func TestAlbumServiceIntegration_GetAlbums(t *testing.T) {
	db := test.SetupIntegrationTest(t)
	
	test.ClearDatabase(db)
	
	albumRepo := repository.NewAlbumRepository(db)
	albumMapper := mappers.NewAlbumMapper()
	albumService := NewAlbumService(albumRepo, *albumMapper)
	
	t.Run("GetAlbums returns empty list when no albums exist", func(t *testing.T) {
		albums, err := albumService.GetAlbums()
		
		assert.NoError(t, err)
		assert.Empty(t, albums)
	})
	
	t.Run("GetAlbums returns albums after creating them", func(t *testing.T) {
		test.ClearDatabase(db)
		
		testAlbums := []models.Album{
			{Name: "Abbey Road"},
			{Name: "Dark Side of the Moon"},
		}
		
		for _, album := range testAlbums {
			result := db.Create(&album)
			assert.NoError(t, result.Error)
		}
		
		albums, err := albumService.GetAlbums()
		
		assert.NoError(t, err)
		assert.Len(t, albums, 2)
		
		albumNames := []string{albums[0].Name, albums[1].Name}
		assert.Contains(t, albumNames, "Abbey Road")
		assert.Contains(t, albumNames, "Dark Side of the Moon")
		
		assert.Greater(t, albums[0].Id, int64(0))
		assert.Greater(t, albums[1].Id, int64(0))
	})
}

func TestAlbumServiceIntegration_CreateAlbum(t *testing.T) {
	db := test.SetupIntegrationTest(t)
	
	test.ClearDatabase(db)
	
	albumRepo := repository.NewAlbumRepository(db)
	albumMapper := mappers.NewAlbumMapper()
	albumService := NewAlbumService(albumRepo, *albumMapper)
	
	t.Run("CreateAlbum successfully creates and returns album", func(t *testing.T) {
		albumName := "Nevermind"
		
		createdAlbum, err := albumService.CreateAlbum(albumName)
		
		assert.NoError(t, err)
		assert.Equal(t, albumName, createdAlbum.Name)
		assert.Greater(t, createdAlbum.Id, int64(0))
		
		var dbAlbum models.Album
		result := db.First(&dbAlbum, createdAlbum.Id)
		assert.NoError(t, result.Error)
		assert.Equal(t, albumName, dbAlbum.Name)
		assert.Equal(t, uint(createdAlbum.Id), dbAlbum.ID)
	})
	
	t.Run("CreateAlbum persists multiple albums", func(t *testing.T) {
		test.ClearDatabase(db)
		
		albums := []string{"Album 1", "Album 2", "Album 3"}
		var createdIDs []int64
		
		for _, name := range albums {
			createdAlbum, err := albumService.CreateAlbum(name)
			assert.NoError(t, err)
			assert.Equal(t, name, createdAlbum.Name)
			createdIDs = append(createdIDs, createdAlbum.Id)
		}
		
		allAlbums, err := albumService.GetAlbums()
		assert.NoError(t, err)
		assert.Len(t, allAlbums, 3)
		
		for _, createdID := range createdIDs {
			found := false
			for _, album := range allAlbums {
				if album.Id == createdID {
					found = true
					break
				}
			}
			assert.True(t, found, "Created album with ID %d should be found in GetAlbums result", createdID)
		}
	})
}
