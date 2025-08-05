package domain

import (
	"time"

	"github.com/google/uuid"
)

type Artist interface {
	ID() uuid.UUID
	Name() string
	ProfileCoverPhotoPath() string
	CreatedAt() time.Time
}

type artist struct {
	id uuid.UUID
	name                  string
	profileCoverPhotoPath string
	createdAt time.Time
}

func (a *artist) ID() uuid.UUID {
	return a.id
}

func (a *artist) Name() string {
	return a.name
}

func (a *artist) ProfileCoverPhotoPath() string {
	return a.profileCoverPhotoPath
}

func (a *artist) CreatedAt() time.Time {
	return a.createdAt
}

func NewArtist(id uuid.UUID, name string, profilePhoto string, createdAt time.Time) Artist {
	return &artist{
		id,
		name,
		profilePhoto,
		createdAt,
	}
}
