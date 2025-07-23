package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/joaopedropio/musiquera/app"
)

func NewReleaseController(a app.Application) *ReleaseController {
	return &ReleaseController{
		application: a,
	}
}

type ReleaseController struct {
	application app.Application
}

type Artist struct {
	Name             string `json:"name"`
	ProfileCoverPath string `json:"profileCoverPath"`
}

type Track struct {
	Name string `json:"name"`
	File string `json:"file"`
}

type ReleaseResponse struct {
	Name        string  `json:"name"`
	Artist      Artist  `json:"artist"`
	Cover       string  `json:"cover"`
	ReleaseDate string  `json:"releaseDate"`
	Tracks       []*Track `json:"tracks"`
}

func (c *ReleaseController) GetReleasesByArtist(w http.ResponseWriter, r *http.Request) {
	artistName := chi.URLParam(r, "artistName")
	if artistName == "" {
		http.Error(w, "artistName cannot be empty", 500)
	}
	fullReleases, err := c.application.Repo().GetReleasesByArtist(artistName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get most recent release: %s", err.Error()), 500)
	}
	var releaseResponses []ReleaseResponse
	for _, release := range fullReleases {
		var tracks []*Track
		for _, track := range release.Tracks() {
			s := &Track{
				Name: track.Name(),
				File: track.File(),
			}
			tracks = append(tracks, s)
		}
		releaseResponse := ReleaseResponse{
			Name:  release.Name(),
			Cover: release.Cover(),
			Artist: Artist{
				release.Artist().Name(),
				release.Artist().ProfileCoverPhotoPath(),
			},
			ReleaseDate: release.ReleaseDate().String(),
			Tracks:       tracks,
		}
		releaseResponses = append(releaseResponses, releaseResponse)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(releaseResponses)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal release response: %s", err.Error()), 500)
	}

}

func (c *ReleaseController) GetMostRecent(w http.ResponseWriter, r *http.Request) {
	fullRelease, err := c.application.Repo().GetMostRecentRelease()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get most recent release: %s", err.Error()), 500)
	}
	var tracks []*Track
	for _, track := range fullRelease.Tracks() {
		s := &Track{
			Name: track.Name(),
			File: track.File(),
		}
		tracks = append(tracks, s)
	}
	release := ReleaseResponse{
		Name: fullRelease.Name(),
		Artist: Artist{
			fullRelease.Artist().Name(),
			fullRelease.Artist().ProfileCoverPhotoPath(),
		},
		ReleaseDate: fullRelease.ReleaseDate().String(),
		Tracks:       tracks,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(release)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal release response: %s", err.Error()), 500)
	}
}

func (c *ReleaseController) Get(w http.ResponseWriter, r *http.Request) {
	releaseID := chi.URLParam(r, "releaseID")
	if releaseID == "" {
		http.Error(w, "releaseID must not be empty", 400)
		return
	}
	id, err := uuid.Parse(releaseID)
	if err != nil {
		http.Error(w, fmt.Sprintf("releaseID must be an uuid: releasID: %s", releaseID), 400)
		return
	}
	release, err := c.application.Repo().GetRelease(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get release by id: id: %s, %s", releaseID, err.Error()), 500)
	}
	_, err = fmt.Fprintf(w, "name: %s, releaseDate: %s", release.Name(), release.ReleaseDate().String())
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to write to response: %s", err.Error()), 500)
	}
}
