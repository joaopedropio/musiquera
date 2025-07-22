package infra

import (
	"fmt"
	"github.com/google/uuid"
	"time"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	repodomain "github.com/joaopedropio/musiquera/app/domain/repo"
)

type repo struct {
	albums []domain.FullAlbum
}

func (r *repo) GetArtists() ([]domain.Artist, error) {
	var artists []domain.Artist
	for _, album := range r.albums {
		artists = append(artists, album.Artist())
	}
	return UniqueBy(artists, func(a domain.Artist) string {
		return a.Name()
	}), nil
}

func (r *repo) GetAlbumsByArtist(artistName string) ([]domain.FullAlbum, error) {
	var albums []domain.FullAlbum
	for _, album := range r.albums {
		if album.Artist().Name() == artistName {
			albums = append(albums, album)
		}
	}
	return albums, nil
}

func (r *repo) AddAlbum(name string, cover string, release domain.Date, artist domain.Artist, songs []domain.Song) (uuid.UUID, error) {
	id := uuid.New()
	a := domain.NewFullAlbum(
		id,
		name,
		cover,
		release,
		artist,
		songs,
		time.Now(),
	)
	r.albums = append(r.albums, a)
	return id, nil
}

func (r *repo) GetAlbum(id uuid.UUID) (domain.Album, error) {
	fullAlbum, err := r.GetFullAlbum(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get full album: %w", err)
	}
	return domain.NewAlbum(fullAlbum.Name(), fullAlbum.ReleaseDate()), nil
}

func (r *repo) GetFullAlbum(id uuid.UUID) (domain.FullAlbum, error) {
	for _, a := range r.albums {
		if a.ID().String() == id.String() {
			//return a.name, a.release, a.artist, a.songs, nil
			return a, nil
		}
	}
	return nil, fmt.Errorf("unable to find album with id %s", id.String())
}

func (r *repo) GetMostRecentAlbum() (domain.FullAlbum, error) {
	var a domain.FullAlbum
	for _, album := range r.albums {
		if a == nil {
			a = album
			continue
		}
		if album.CreatedAt().After(a.CreatedAt()) {
			a = album
		}
	}
	if a == nil {
		return nil, fmt.Errorf("no album on database")
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
