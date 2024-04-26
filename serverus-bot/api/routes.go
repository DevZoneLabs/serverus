package api

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux();

	mux.HandleFunc("GET /healthcheck", s.healthCheck)

	//mux.HandleFunc ("POST /sendMessage" ,s.sendMessage)

	return mux
}