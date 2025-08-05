package domain

import (
	"github.com/google/uuid"
	"time"
)

const ReleaseTypeAlbum = ReleaseType("album")
const ReleaseTypeLiveSet = ReleaseType("liveSet")

type ReleaseType string

type fullRelease struct {
	id        uuid.UUID
	name      string
	releaseType ReleaseType
	cover     string
	release   Date
	artist    Artist
	tracks     []Track
	createdAt time.Time
}

func NewFullRelease(id uuid.UUID, name string, releaseType ReleaseType, cover string, releaseDate Date, artist Artist, tracks []Track, createdAt time.Time) FullRelease {
	return &fullRelease{
		id:        id,
		name:      name,
		cover:     cover,
		releaseType: releaseType,
		release:   releaseDate,
		artist:    artist,
		tracks:     tracks,
		createdAt: createdAt,
	}
}

type FullRelease interface {
	ID() uuid.UUID
	Name() string
	Cover() string
	Type() ReleaseType
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

func (a *fullRelease) Type() ReleaseType {
	return a.releaseType
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
