package main

import (
	"net/http"
)

func (server *Config) routes() http.Handler {

	router := http.NewServeMux()

	// TO-DO
	// ADD healthcheck handler

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	// TO-DO
	// Implement healthcheck logic
}