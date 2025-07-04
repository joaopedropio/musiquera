package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	repodomain "github.com/joaopedropio/musiquera/app/domain/repo"
)

type album struct {
	id        uuid.UUID
	name      string
	release   domain.Date
	artist    domain.Artist
	songs     []domain.Song
	createdAt time.Time
}

type repo struct {
	albums []*album
}

func (r *repo) AddAlbum(name string, release domain.Date, artist domain.Artist, songs []domain.Song) (uuid.UUID, error) {
	id := uuid.New()
	a := &album{
		id,
		name,
		release,
		artist,
		songs,
		time.Now(),
	}
	r.albums = append(r.albums, a)
	return id, nil
}

func (r *repo) GetAlbum(id uuid.UUID) (domain.Album, error) {
	name, releaseDate, _, _, err := r.GetFullAlbum(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get full album: %w", err)
	}
	return domain.NewAlbum(name, releaseDate), nil
}

func (r *repo) GetFullAlbum(id uuid.UUID) (string, domain.Date, domain.Artist, []domain.Song, error) {
	for _, a := range r.albums {
		if a.id.String() == id.String() {
			return a.name, a.release, a.artist, a.songs, nil
		}
	}
	return "", nil, nil, nil, fmt.Errorf("unable to find album with id %s", id.String())
}

func (r *repo) GetMostRecentAlbum() (string, domain.Date, domain.Artist, []domain.Song, error) {
	var a *album
	for _, album := range r.albums {
		if a == nil {
			a = album
			continue
		}
		if album.createdAt.After(a.createdAt) {
			a = album
		}
	}
	if a == nil {
		return "", nil, nil, nil, fmt.Errorf("no album on database")
	}
	return a.name, a.release, a.artist, a.songs, nil
}

func NewRepo() repodomain.Repo {
	return &repo{}
}
