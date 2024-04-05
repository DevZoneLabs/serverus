package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)
type Bot struct {
	Session *discordgo.Session
}

func Run(token string) (error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("error establishing Discord session: %v", err)
	}
	
	bot := &Bot{
		Session: session,
	}

	// Register Bot Handlers
	bot.RegisterHandlers()

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		return fmt.Errorf("error opening connection to Discord: %v", err)
	}
	defer session.Close()

	fmt.Println("Serverus-Bot online!")

	// Create a channel to keep this function running until it terminates.
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-channel

	return nil
}