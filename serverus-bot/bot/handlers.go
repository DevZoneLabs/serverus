package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) RegisterHandlers() {
	b.Session.AddHandler(b.helloWorld)
}

func (b *Bot) helloWorld(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return 
	}

	if message.Content == "hello" {
		session.ChannelMessageSend(message.ChannelID, "World!")
	}
}