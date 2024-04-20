package bot

import (
	"log"
	"context"

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

	bot := &Bot {
		Session: session,
	}

	return bot, nil
}

func (bot *Bot) Run(ctx context.Context) {

	err := bot.Session.Open()
	if err != nil {
		log.Panic(err)
	}
	defer bot.Session.Close()

	log.Println("Serverus is online")
	
	<- ctx.Done()

	if err := bot.Session.Close(); err != nil {
		log.Fatal("Error shutting down the Bot: ", err)
	}

	log.Println("Bot Stopped")
}

