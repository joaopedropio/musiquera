package infra_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
)

func TestRepo_ShouldAddAlbum_WhenAlbumIsAdded(t *testing.T) {
	// Arrange
	repo := infra.NewRepo()
	name := "album_name"
	artist := domain.NewArtist("artist_name")
	release := domain.NewDate(2000, 1, 1)
	songs := []domain.Song{}

	// Act
	id, err := repo.AddAlbum(name, release, artist, songs)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, uuid.Validate(id.String()))
}
