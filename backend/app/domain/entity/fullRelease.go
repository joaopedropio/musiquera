package domain

import (
	"github.com/google/uuid"
	"time"
)

type fullRelease struct {
	id        uuid.UUID
	name      string
	cover     string
	release   Date
	artist    Artist
	tracks     []Track
	createdAt time.Time
}

func NewFullRelease(id uuid.UUID, name string, cover string, release Date, artist Artist, tracks []Track, createdAt time.Time) FullRelease {
	return &fullRelease{
		id:        id,
		name:      name,
		cover:     cover,
		release:   release,
		artist:    artist,
		tracks:     tracks,
		createdAt: createdAt,
	}
}

type FullRelease interface {
	ID() uuid.UUID
	Name() string
	Cover() string
	ReleaseDate() Date
	Artist() Artist
	Tracks() []Track
	CreatedAt() time.Time
}

func (a *fullRelease) ID() uuid.UUID {
	return a.id
}

func (a *fullRelease) Name() string {
	return a.name
}

func (a *fullRelease) Cover() string {
	return a.cover
}
func (a *fullRelease) ReleaseDate() Date {
	return a.release
}

func (a *fullRelease) Artist() Artist {
	return a.artist
}

func (a *fullRelease) Tracks() []Track {
	return a.tracks
}

func (a *fullRelease) CreatedAt() time.Time {
	return a.createdAt
}
