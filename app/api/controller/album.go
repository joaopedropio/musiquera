package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

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

func (c *AlbumController) Get(w http.ResponseWriter, r *http.Request) {
	albumID := chi.URLParam(r, "albumID")
	if albumID == "" {
		http.Error(w, fmt.Sprintf("albumID must not be empty"), 400)
		return
	}
	id, err := uuid.Parse(albumID)
	if err != nil {
		http.Error(w, fmt.Sprintf("albumID must be an uuid: albumID: %s", albumID), 400)
		return
	}
	album, err := c.application.Repo().GetAlbum(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get album by id: id: %s, %w", albumID, err), 500)
	}
	w.Write([]byte(fmt.Sprintf("name: %s, releaseDate: %s", album.Name(), album.ReleaseDate().String())))
}
