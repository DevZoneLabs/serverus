package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"serverus-bot/api"
	"serverus-bot/bot"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
)

const (
	portNum = ":80"
	gracePeriod = 15
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	log.Println(os.Getenv("SERVERUS_TOKEN"))
	serverus := bot.NewBot(os.Getenv("SERVERUS_TOKEN"))

	server := api.NewServer(portNum, serverus)

	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(){
		defer wg.Done()
		serverus.Run(ctx)
	}()

	go server.Start()

	// Create a channel to listen for interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<- interrupt

	// Shutdown the Server and the Bot
	server.Shutdown(gracePeriod)
	cancel()
	wg.Wait()
}