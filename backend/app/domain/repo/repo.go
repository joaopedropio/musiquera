package repo

import (
	"github.com/google/uuid"
	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type Repo interface {
	GetReleasesByArtist(artistName string) ([]entity.FullRelease, error)
	GetArtists() ([]entity.Artist, error)
	AddRelease(name string, cover string, release entity.Date, artist entity.Artist, tracks []entity.Track) (uuid.UUID, error)
	GetRelease(id uuid.UUID) (entity.Release, error)
	GetFullRelease(id uuid.UUID) (entity.FullRelease, error)
	GetMostRecentRelease() (entity.FullRelease, error)
}
