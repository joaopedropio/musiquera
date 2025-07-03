package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api"
)

func main() {
	if err := runApi(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func runApi() error {
	r := chi.NewRouter()
	a, err := app.NewApplication()
	api.ConfigureAPI(r, a)
	if err != nil {
		return fmt.Errorf("unable to configure api")
	}
	port := ":8080"
	fmt.Print("musiquera running on port ", port)
	return http.ListenAndServe(port, r)
}
