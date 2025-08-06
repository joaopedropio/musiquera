package infra_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/database"
	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
	"github.com/joaopedropio/musiquera/app/utils"
)

func TestRepo_AddArtist(t *testing.T) {
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewRepo(db)
	id := uuid.New()
	name := "Artist Name"
	cover := "/some/path"
	createdAt := time.Now()
	artist := domain.NewArtist(id, name, cover, createdAt)
	err := repo.AddArtist(artist)

	assert.Nil(t, err)

	artists, err := repo.GetArtists()
	assert.Nil(t, err)

	a, ok := utils.Find(artists, func(a domain.Artist) bool {
		return a.Name() == name
	})
	assert.True(t, ok)
	assert.Equal(t, id, a.ID())
	assert.Equal(t, name, a.Name())
	assert.Equal(t, cover, a.ProfileCoverPhotoPath())
	assert.True(t, utils.IsTimeEqual(createdAt, a.CreatedAt()))
}
