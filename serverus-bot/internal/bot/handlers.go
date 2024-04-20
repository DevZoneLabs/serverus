package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const SERVER_DOWN_MESSAGE = "Sorry! The Server is not taking requests at this time!"

func (bot *Bot) RegisterHandlers() {
	// Registers handler to the bot
	bot.AddEventHandler(helloWorld)
}

func (bot *Bot) AddEventHandler(handler interface {}) func() {
	// Customizes the native AddHandler function making us able to add additional logic to the pre-processing of the handler.
	switch h := handler.(type) {
		case func(*discordgo.Session, *discordgo.MessageCreate): 
			return bot.Session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {

				if message.Author.ID == session.State.User.ID {
					return
				}
				if !bot.AcceptingRequests {
					session.ChannelMessageSend(message.ChannelID, SERVER_DOWN_MESSAGE)
					return
				}
				h(session, message)
			})
		default:
			log.Println("Invalid Handler Function")
			return nil
	}
}

// Handlers

func helloWorld(session *discordgo.Session, message *discordgo.MessageCreate){
	if message.Author.ID == session.State.User.ID {
		return 
	}

	if message.Content == "Hello" {
		session.ChannelMessageSend(message.ChannelID, "World!")
	}
}
