package api

import (
	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api/controller"
)

type Controller struct {
	Album    *controller.AlbumController
	PingPong *controller.PingPongController
	StaticFiles *controller.StaticController
}

func SetupControllers(a app.Application) Controller {
	albumController := controller.NewAlbumController(a)
	pingpingController := controller.NewPingPong()
	staticFilesController := controller.NewStaticController(a.Environment().WebStaticFilesDir)
	return Controller{
		Album:    albumController,
		PingPong: pingpingController,
		StaticFiles: staticFilesController,
	}
}
