package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) routes() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/healthcheck", s.healthCheck)

	return mux
}