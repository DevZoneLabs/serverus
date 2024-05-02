package api

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", s.healthCheck)
	mux.HandleFunc("POST /v1/sendchannelmessage", s.sendChannelMessage)

	return mux
}
