package repo

import (
	"github.com/google/uuid"
	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type Repo interface {
	AddArtist(artist entity.Artist) error
	GetReleasesByArtist(artistName string) ([]entity.FullRelease, error)
	GetArtists() ([]entity.Artist, error)
	AddFullRelease(fullRelease entity.FullRelease) error
	GetRelease(id uuid.UUID) (entity.Release, error)
	GetFullRelease(id uuid.UUID) (entity.FullRelease, error)
	GetMostRecentRelease() (entity.FullRelease, error)
}
