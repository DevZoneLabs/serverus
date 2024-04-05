package main

import (
	"log"
	"serverus-bot/bot"
	"serverus-bot/config"
	"sync"
)

// This is the entry point of our service

func main(){
	
	// Load Configuration File
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// synchronizes the main routine with goroutines. We increment the counter of the WaitGroup before starting each goroutine
	// and decrement it when every go routine finishes
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	// Initialize Bot
	go func () {
		defer waitGroup.Done()
		if err := bot.Run(conf.DiscordToken); err != nil {
			log.Fatal(err)
		}
	}()
		

	// TO-DO 
	// Initialize server

	waitGroup.Wait()
}