package api

import (
	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api/controller"
)

type Controller struct {
	Album *controller.AlbumController
}

func SetupControllers(a app.Application) Controller {
	albumController := controller.NewAlbumController(a)
	return Controller{
		Album: albumController,
	}
}
