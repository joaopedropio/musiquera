package api

import (
	chi "github.com/go-chi/chi/v5"

	app "github.com/joaopedropio/musiquera/app"
)

func ConfigureAPI(r *chi.Mux, a app.Application) error {
	c := SetupControllers(a)
	configureHandlers(r, c)
	return nil
}

func configureHandlers(r *chi.Mux, c Controller) {
	r.Get("/api/album/{albumID}", c.Album.Get)
}
