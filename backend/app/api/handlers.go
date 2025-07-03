package api

import (
	"github.com/go-chi/chi/v5"
)

func ConfigureHandlers(r *chi.Mux, c Controller) {
	r.Get("/api/album/{albumID}", c.Album.Get)
}
