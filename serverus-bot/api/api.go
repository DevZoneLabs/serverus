package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serverus-bot/bot"
	"time"
)

const werPort = "80"

type Config struct {
	Bot *bot.Bot
}

func Initialize(app *Config) {
    log.Println("Starting api server!")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", werPort),
        Handler: app.routes(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v\n", err)
		}
	}()

	// Handle shutdown
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop

		log.Println("Shutting down API server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v\n", err)
		}

		log.Println("API server gracefully stopped")
	}()
	
}