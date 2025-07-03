package repo

import (
	"github.com/google/uuid"
	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type Repo interface {
	AddAlbum(name string, release entity.Date, artist entity.Artist, songs []entity.Song) (uuid.UUID, error)
	GetAlbum(id uuid.UUID) (entity.Album, error)
	GetFullAlbum(id uuid.UUID) (string, entity.Date, entity.Artist, []entity.Song, error)
}
