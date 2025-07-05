package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/joaopedropio/musiquera/app"
)

func NewAlbumController(a app.Application) *AlbumController {
	return &AlbumController{
		application: a,
	}
}

type AlbumController struct {
	application app.Application
}

type Song struct {
	Name string `json:"name"`
	File string `json:"file"`
}

type AlbumResponse struct {
	Name        string  `json:"name"`
	Artist      string  `json:"artist"`
	ReleaseDate string  `json:"releaseDate"`
	Songs       []*Song `json:"songs"`
}

func (c *AlbumController) GetAlbumsByArtist(w http.ResponseWriter, r *http.Request) {
	artistName := chi.URLParam(r, "artistName")
	if artistName == "" {
		http.Error(w, "artistName cannot be empty", 500)
	}
	fullAlbums, err := c.application.Repo().GetAlbumsByArtist(artistName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get most recent album: %s", err.Error()), 500)
	}
	var albumResponses []AlbumResponse
	for _, album := range fullAlbums {
		var songs []*Song
		for _, song := range album.Songs() {
			s := &Song{
				Name: song.Name(),
				File: song.File(),
			}
			songs = append(songs, s)
		}
		albumResponse := AlbumResponse{
			Name:        album.Name(),
			Artist:      album.Artist().Name(),
			ReleaseDate: album.ReleaseDate().String(),
			Songs:       songs,
		}
		albumResponses = append(albumResponses, albumResponse)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(albumResponses)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal album response: %s", err.Error()), 500)
	}

}

func (c *AlbumController) GetMostRecent(w http.ResponseWriter, r *http.Request) {
	fullAlbum, err := c.application.Repo().GetMostRecentAlbum()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get most recent album: %s", err.Error()), 500)
	}
	var songss []*Song
	for _, song := range fullAlbum.Songs() {
		s := &Song{
			Name: song.Name(),
			File: song.File(),
		}
		songss = append(songss, s)
	}
	album := AlbumResponse{
		Name:        fullAlbum.Name(),
		Artist:      fullAlbum.Artist().Name(),
		ReleaseDate: fullAlbum.ReleaseDate().String(),
		Songs:       songss,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(album)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal album response: %s", err.Error()), 500)
	}
}

func (c *AlbumController) Get(w http.ResponseWriter, r *http.Request) {
	albumID := chi.URLParam(r, "albumID")
	if albumID == "" {
		http.Error(w, "albumID must not be empty", 400)
		return
	}
	id, err := uuid.Parse(albumID)
	if err != nil {
		http.Error(w, fmt.Sprintf("albumID must be an uuid: albumID: %s", albumID), 400)
		return
	}
	album, err := c.application.Repo().GetAlbum(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get album by id: id: %s, %s", albumID, err.Error()), 500)
	}
	_, err = fmt.Fprintf(w, "name: %s, releaseDate: %s", album.Name(), album.ReleaseDate().String())
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to write to response: %s", err.Error()), 500)
	}
}
