package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
}

func Init(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	bot := &Bot{
		Session: session,
	}

	return bot, nil
}

func (bot *Bot) Run() {

	session := bot.Session

	bot.RegisterHandlers()

	err := session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Println("Serverus is online!")

	// Create a channel to keep this function running until it terminates.
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-channel
}
