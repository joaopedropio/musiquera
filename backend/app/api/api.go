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
	r.Get("/api/artist/", c.Artist.GetAllArtists)
	r.Get("/api/release/{releaseID}", c.Release.Get)
	r.Get("/api/release/mostRecent", c.Release.GetMostRecent)
	r.Get("/api/release/byArtist/{artistName}", c.Release.GetReleasesByArtist)
	r.Get("/ping", c.PingPong.Get)
	r.NotFound(c.StaticFiles.ServeStatic)
}
