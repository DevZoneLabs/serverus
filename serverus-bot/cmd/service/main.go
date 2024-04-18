package main

import (
	"log"
	"serverus-bot/api"
	"serverus-bot/bot"
	"serverus-bot/config"
	"sync"
)

// This is the entry point of our service
func main() {
	// Load Configuration File
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Bot
	b, err := bot.Init(conf.DiscordToken)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize API server
	api.Initialize(&api.Config{
		Bot: b,
	})

	// Start bot
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		b.Run() // Run Bot
	}()

	waitGroup.Wait()
}