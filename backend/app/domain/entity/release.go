package domain

import (
	"github.com/google/uuid"
	"time"
)

type Release interface {
	ID() uuid.UUID
	Name() string
	Cover() string
	ReleaseDate() Date
	Artist() Artist
	CreatedAt() time.Time
}

type release struct {
	id          uuid.UUID
	name        string
	releaseDate Date
	cover       string
	artist      Artist
	createdAt   time.Time
}

func (a *release) ID() uuid.UUID {
	return a.id
}

func (a *release) Name() string {
	return a.name
}

func (a *release) Cover() string {
	return a.cover
}

func (a *release) ReleaseDate() Date {
	return a.releaseDate
}

func (a *release) Artist() Artist {
	return a.artist
}

func (a *release) CreatedAt() time.Time {
	return a.createdAt
}

func NewRelease(name string, releaseDate Date, cover string, artist Artist) Release {
	return &release{
		uuid.New(),
		name,
		releaseDate,
		cover,
		artist,
		time.Now(),
	}
}
