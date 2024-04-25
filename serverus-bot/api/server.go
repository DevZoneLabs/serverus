package api

import (
	"context"
	"log"
	"net/http"
	"serverus-bot/bot"
	"time"
)

type Server struct {
	listenAddrs string
	bot bot.Bot
	srv *http.Server
}

func NewServer(listenAddrs string, bot *bot.Bot) *Server {
	return &Server{
		listenAddrs: listenAddrs,
		bot: *bot,
	}
}

func (s *Server) Start() error {

	log.Println("Starting server...")
	
	srv := &http.Server{
		Addr: s.listenAddrs,
		Handler: s.routes(),
	}

	s.srv = srv

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown(gracePeriod int) error {

	s.bot.StopAcceptingRequests()

	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second * time.Duration(gracePeriod))
	defer cancelTimeout()

	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Println("Error shutting down the server: ", err)
		return err
	}

	log.Println("Server stopped!")
	return nil
}