package bot

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const heartbeatTimeout = 30 * time.Minute

type Bot struct {
	session           *discordgo.Session
	acceptingRequests bool
}

func NewBot(botToken string) *Bot {

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic(err)
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	return &Bot{
		acceptingRequests: true,
		session:           session,
	}
}

func (bot *Bot) Run(ctx context.Context) {

	log.Print("Initializing bot...")

	// register bot handlers
	bot.registerHandlers()

	err := bot.session.Open()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Bot is now online!")

	go bot.monitorHeartbeats()

	<-ctx.Done()

	if err := bot.session.Close(); err != nil {
		log.Panicf("Error shutting down the Bot: ", err)
	}

	log.Println("Bot stopped")
}

func (bot *Bot) Close() error {
	return bot.session.Close()
}

func (bot *Bot) StopAcceptingRequests() {
	bot.acceptingRequests = false
}


func (bot *Bot) monitorHeartbeats() {
	ticker := time.NewTicker(heartbeatTimeout)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			// Check if the heartbeat is received
			if !bot.session.DataReady {
				log.Println("No heartbeat ACK received, reconnecting...")
				err := b.reconnect()
				if err != nil {
					log.Printf("Bot failed to reconnect: %v", err)
				}
			}
		}
	}
}

func (bot *Bot) reconnect() error {
	err := bot.session.Close()
	if err != nil {
		log.Printf("Error closing session during reconnect: %v", err)
	}
	time.Sleep(5 * time.Second)
	err = bot.session.Open()
	if err != nil {
		return err
	}

	log.Println("Reconnected to Discord")
	return nil

}