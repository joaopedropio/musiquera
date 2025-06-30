package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api"
)

func main() {
	r := chi.NewRouter()
	a, err := app.NewApplication()
	api.ConfigureAPI(r, a)
	if err != nil {
		panic(err)
	}
	port := ":8080"
	fmt.Print("musiquera running on port ", port)
	http.ListenAndServe(port, r)
}
