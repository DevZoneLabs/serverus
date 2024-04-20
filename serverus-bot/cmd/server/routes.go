package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (server *Config) routes() *chi.Mux {

	mux := chi.NewMux()

	mux.Use(middleware.Heartbeat("/ping"))

	return mux
}
