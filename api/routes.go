package api

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO - Healtcheck should either be a get request or a post to a channel. currently broken!
	mux.HandleFunc("GET /healthcheck", s.healthCheck)
	mux.HandleFunc("POST /v1/sendchannelmessage", s.sendChannelMessage)

	return mux
}
