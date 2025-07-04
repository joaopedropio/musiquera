package api

import (
	chi "github.com/go-chi/chi/v5"
	app "github.com/joaopedropio/musiquera/app"
)

func ConfigureAPI(r *chi.Mux, a app.Application) {
	c := SetupControllers(a)
	configureHandlers(r, c)
}

func configureHandlers(r *chi.Mux, c Controller) {
	r.Get("/api/album/{albumID}", c.Album.Get)
	r.Get("/api/album/mostRecent", c.Album.GetMostRecent)
	r.Get("/ping", c.PingPong.Get)
	r.NotFound(c.StaticFiles.ServeStatic)
}
