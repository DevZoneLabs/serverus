package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serverus-bot/internal/bot"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Bot bot.Bot
}

const portNum = ":80"

func main() {

	// Read .env file
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	// Initialize the bot using the SERVERUS_TOKEN from the environment
	bot, err := bot.Init(os.Getenv("SERVERUS_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	// Initialize app configuration
	app := Config {
		Bot: *bot,
	}

	// Create a server object
	srv := &http.Server{
		Addr: portNum,
		Handler: app.routes(),
	}

	// Create a context to handle shutting down the bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a waitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the Bot in a separate goroutine and pass the context
	go func() {
		defer wg.Done()
		bot.Run(ctx)
	}()

	// Run the Server in a separate goroutine
	go func() {
		log.Println("Server listening on port 80")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Create a channel to listen for interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<- interrupt
	
	// Handle graceful shutdown
	log.Println("Shutting down...")

	// Prevent the Bot from Accepting any requests
	bot.SetAcceptingRequests(false)

	log.Println("Bot Not Taking New Requests")

	ctx, cancelTimeOut := context.WithTimeout(ctx, 15 * time.Second)
	defer cancelTimeOut()

	// Trigger a server shutdown with the same context
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")

	// Stop the Bot
	cancel()

	wg.Wait()

	log.Println("Serverus-Bot Service Offline")
}