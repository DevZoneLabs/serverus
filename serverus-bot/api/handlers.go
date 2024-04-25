package api

import (
	"log"
	"net/http"
)

func (s *Server) healthCheck(w http.ResponseWriter, _ *http.Request) {

	err := s.bot.HealthCheckMessage("750064736351027228")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		w.WriteHeader(http.StatusOK)
	}
	
}