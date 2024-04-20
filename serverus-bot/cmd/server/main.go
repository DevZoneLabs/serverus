package main

import (
	"log"
	"net/http"
	"os"
	"serverus-bot/internal/bot"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Bot bot.Bot
}

const portNum = ":80"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	bot, err := bot.Init(os.Getenv("SERVERUS_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	app := Config {
		Bot: *bot,
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		bot.Run()
	}()

	srv := &http.Server{
		Addr: portNum,
		// TO-DO Add Handler: app.routes()
	}

	go func() {
		log.Println("Server listening on port 80")
		srv.ListenAndServe()
	}()

	waitGroup.Wait()
	
}