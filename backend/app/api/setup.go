package api

import (
	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api/controller"
)

type Controller struct {
	Album    *controller.AlbumController
	Artist *controller.ArtistController
	PingPong *controller.PingPongController
	StaticFiles *controller.StaticController
}

func SetupControllers(a app.Application) Controller {
	artistController := controller.NewArtistController(a)
	albumController := controller.NewAlbumController(a)
	pingpingController := controller.NewPingPong()
	staticFilesController := controller.NewStaticController(a.Environment().WebStaticFilesDir)
	return Controller{
		Artist: artistController,
		Album:    albumController,
		PingPong: pingpingController,
		StaticFiles: staticFilesController,
	}
}
