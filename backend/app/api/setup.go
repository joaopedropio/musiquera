package api

import (
	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api/controller"
)

type Controller struct {
	Release     *controller.ReleaseController
	Artist      *controller.ArtistController
	PingPong    *controller.PingPongController
	StaticFiles *controller.StaticController
	Security    *controller.SecurityController
}

func SetupControllers(a app.Application) Controller {
	artistController := controller.NewArtistController(a)
	releaseController := controller.NewReleaseController(a)
	pingpingController := controller.NewPingPong()
	staticFilesController := controller.NewStaticController(a.Environment().WebStaticFilesDir)
	securityController := controller.NewSecurityController(a)
	return Controller{
		Artist:      artistController,
		Release:     releaseController,
		PingPong:    pingpingController,
		StaticFiles: staticFilesController,
		Security:    securityController,
	}
}
