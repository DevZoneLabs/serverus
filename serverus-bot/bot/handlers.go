package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const SERVER_DOWN_MESSAGE = "Sorry! I cannot your request right now!"

func (bot *Bot) registerHandlers() {

	bot.addHandler(helloWorld)

}

func (bot *Bot) addHandler(handler interface{}) func() {
	switch h := handler.(type) {
	case func(*discordgo.Session, *discordgo.MessageCreate):
		return bot.session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
			if message.Author.ID == session.State.User.ID {
				return
			}

			if !bot.acceptingRequests {
				session.ChannelMessageSend(message.ChannelID, SERVER_DOWN_MESSAGE)
				return
			}

			h(session, message)
		})
	default:
		log.Panic("Invalid Handler Function")
		return nil
	}
}

func helloWorld(session *discordgo.Session, message *discordgo.MessageCreate){
	
	if message.Content == "Hello" {
		session.ChannelMessageSend(message.ChannelID, "World!")
	}
}