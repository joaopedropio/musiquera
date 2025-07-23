package infra

import (
	"fmt"
	"github.com/google/uuid"
	"time"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	repodomain "github.com/joaopedropio/musiquera/app/domain/repo"
)

type repo struct {
	releases []domain.FullRelease
}

func (r *repo) GetArtists() ([]domain.Artist, error) {
	var artists []domain.Artist
	for _, release := range r.releases {
		artists = append(artists, release.Artist())
	}
	return UniqueBy(artists, func(a domain.Artist) string {
		return a.Name()
	}), nil
}

func (r *repo) GetReleasesByArtist(artistName string) ([]domain.FullRelease, error) {
	var releases []domain.FullRelease
	for _, release := range r.releases {
		if release.Artist().Name() == artistName {
			releases = append(releases, release)
		}
	}
	return releases, nil
}

func (r *repo) AddRelease(name string, cover string, release domain.Date, artist domain.Artist, songs []domain.Song) (uuid.UUID, error) {
	id := uuid.New()
	a := domain.NewFullRelease(
		id,
		name,
		cover,
		release,
		artist,
		songs,
		time.Now(),
	)
	r.releases = append(r.releases, a)
	return id, nil
}

func (r *repo) GetRelease(id uuid.UUID) (domain.Release, error) {
	fullRelease, err := r.GetFullRelease(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get full release: %w", err)
	}
	return domain.NewRelease(
		fullRelease.Name(),
		fullRelease.ReleaseDate(),
		fullRelease.Cover(),
		fullRelease.Artist(),
	), nil
}

func (r *repo) GetFullRelease(id uuid.UUID) (domain.FullRelease, error) {
	for _, a := range r.releases {
		if a.ID().String() == id.String() {
			//return a.name, a.release, a.artist, a.songs, nil
			return a, nil
		}
	}
	return nil, fmt.Errorf("unable to find release with id %s", id.String())
}

func (r *repo) GetMostRecentRelease() (domain.FullRelease, error) {
	var a domain.FullRelease
	for _, release := range r.releases {
		if a == nil {
			a = release
			continue
		}
		if release.CreatedAt().After(a.CreatedAt()) {
			a = release
		}
	}
	if a == nil {
		return nil, fmt.Errorf("no release on database")
	}
	return a, nil
}

func NewRepo() repodomain.Repo {
	return &repo{}
}

func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := input[:0]
	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func UniqueBy[T any, K comparable](input []T, keySelector func(T) K) []T {
	seen := make(map[K]struct{})
	result := input[:0]

	for _, v := range input {
		key := keySelector(v)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
