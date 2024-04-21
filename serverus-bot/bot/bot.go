package bot

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
	acceptingRequests bool
}

func NewBot(botToken string) *Bot{

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic(err)
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	return &Bot{
		acceptingRequests: true,
		session: session,
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

	<- ctx.Done()

	if err := bot.session.Close(); err != nil {
		log.Panic("Error shutting down the Bot: ", err)
	}

	log.Println("Bot stopped")
}

func (bot *Bot) Close() error {
	return bot.session.Close()
}

func (bot *Bot) StopAcceptingRequests() {
	bot.acceptingRequests = false
}