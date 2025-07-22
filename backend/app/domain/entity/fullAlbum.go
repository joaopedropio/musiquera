package domain

import (
	"github.com/google/uuid"
	"time"
)

type fullAlbum struct {
	id        uuid.UUID
	name      string
	cover     string
	release   Date
	artist    Artist
	songs     []Song
	createdAt time.Time
}

func NewFullAlbum(id uuid.UUID, name string, cover string, release Date, artist Artist, songs []Song, createdAt time.Time) FullAlbum {
	return &fullAlbum{
		id:        id,
		name:      name,
		cover:     cover,
		release:   release,
		artist:    artist,
		songs:     songs,
		createdAt: createdAt,
	}
}

type FullAlbum interface {
	ID() uuid.UUID
	Name() string
	Cover() string
	ReleaseDate() Date
	Artist() Artist
	Songs() []Song
	CreatedAt() time.Time
}

func (a *fullAlbum) ID() uuid.UUID {
	return a.id
}

func (a *fullAlbum) Name() string {
	return a.name
}

func (a *fullAlbum) Cover() string {
	return a.cover
}
func (a *fullAlbum) ReleaseDate() Date {
	return a.release
}

func (a *fullAlbum) Artist() Artist {
	return a.artist
}

func (a *fullAlbum) Songs() []Song {
	return a.songs
}

func (a *fullAlbum) CreatedAt() time.Time {
	return a.createdAt
}
