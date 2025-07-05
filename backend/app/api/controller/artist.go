package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaopedropio/musiquera/app"
)

type ArtistResponse struct {
	Name string `json:"name"`
}

type ArtistController struct {
	a app.Application
}

func NewArtistController(a app.Application) *ArtistController {
	return &ArtistController{
		a,
	}
}

func (c *ArtistController) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := c.a.Repo().GetArtists()
	if err!= nil {
		http.Error(w, fmt.Sprintf("unable to get artists: %s",err.Error()),500)
	}
	var artistsResponse []ArtistResponse
	for _, artist := range artists {
		artistsResponse = append(artistsResponse, ArtistResponse{
			Name: artist.Name(),
		})
	}
	err = json.NewEncoder(w).Encode(artistsResponse)
	if err!= nil {
		http.Error(w, fmt.Sprintf("unable to encode json: %s", err.Error()), 500)
	}
}
