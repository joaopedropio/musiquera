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

func (c *AlbumController) GetMostRecent(w http.ResponseWriter, r *http.Request) {
	name, releaseDate, artist, songs, err := c.application.Repo().GetMostRecentAlbum()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get most recent album: %s", err.Error()), 500)
	}
	var songss []*Song
	for _, song := range songs {
		s := &Song{
			Name: song.Name(),
			File: song.File(),
		}
		songss = append(songss, s)
	}
	album := AlbumResponse{
		Name:        name,
		Artist:      artist.Name(),
		ReleaseDate: releaseDate.String(),
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
