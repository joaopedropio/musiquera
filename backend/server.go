package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/api"
)

type Server struct {
	httpServer *http.Server
	a          app.Application
}

func NewServer() (*Server, error) {
	r := chi.NewRouter()
	a, err := app.NewApplication()
	if err != nil {
		return nil, fmt.Errorf("unable to create new application: %w", err)
	}
	api.ConfigureAPI(r, a)
	return &Server{
		httpServer: &http.Server{
			Addr:    a.Environment().HttpPort,
			Handler: r,
		},
		a: a,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Server listening on", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("unable to stop server: %s", err)
	}

	if err := s.a.Close(); err != nil {
		log.Fatalf("unable to close application: %s", err)
	}

	return nil
}
