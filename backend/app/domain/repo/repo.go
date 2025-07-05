package repo

import (
	"github.com/google/uuid"
	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type Repo interface {
	GetAlbumsByArtist(artistName string) ([]entity.FullAlbum, error)
	GetArtists() ([]entity.Artist, error)
	AddAlbum(name string, release entity.Date, artist entity.Artist, songs []entity.Song) (uuid.UUID, error)
	GetAlbum(id uuid.UUID) (entity.Album, error)
	GetFullAlbum(id uuid.UUID) (entity.FullAlbum, error)
	GetMostRecentAlbum() (entity.FullAlbum, error)
}
