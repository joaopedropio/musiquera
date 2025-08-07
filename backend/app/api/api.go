package api

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	jwtauth "github.com/go-chi/jwtauth/v5"

	app "github.com/joaopedropio/musiquera/app"
)

func ConfigureAPI(m *chi.Mux, a app.Application) {
	c := SetupControllers(a)
	configureHandlers(m, c, a)
}

func configureHandlers(m *chi.Mux, c Controller, a app.Application) {
	m.Group(adminRoutes(c, a))
	m.Group(privateRoutes(c, a))
	m.Group(publicRoutes(c))
}

func adminRoutes(c Controller, a app.Application) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(jwtauth.Verifier(a.LoginService().JWTAuth()))
		r.Use(jwtRedirectMiddleware)
		r.Post("/admin/invite/create", c.User.CreateInvite)
	}
}

func privateRoutes(c Controller, a app.Application) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(jwtauth.Verifier(a.LoginService().JWTAuth()))
		r.Use(jwtRedirectMiddleware)
		r.Post("/logout", c.Security.Logout)
		r.Get("/api/artist/", c.Artist.GetAllArtists)
		r.Get("/api/release/{releaseID}", c.Release.Get)
		r.Get("/api/release/mostRecent", c.Release.GetMostRecent)
		r.Get("/api/release/byArtist/{artistName}", c.Release.GetReleasesByArtist)
		r.NotFound(c.StaticFiles.ServeStatic)
	}
}

func publicRoutes(c Controller) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/auth/check", c.Security.AuthCheck)
		r.Get("/auth/checkAdmin", c.Security.AuthCheckAdmin)
		r.Post("/login", c.Security.Login)
		r.Get("/ping", c.PingPong.Get)
		r.NotFound(c.StaticFiles.ServeStatic)
	}
}

func jwtRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())

		// Token is not present or invalid — redirect
		if err != nil || claims == nil {
			http.Redirect(w, r, "/loginPage", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
