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

	bot := &Bot {
		Session: session,
	}

	return bot, nil
}

func (bot *Bot) Run() {

	err := bot.Session.Open()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Serverus is online")
	defer bot.Session.Close()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<- channel
	log.Println("Interrupted!")
}