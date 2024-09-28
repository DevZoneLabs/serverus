package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

const SERVER_DOWN_MESSAGE = "Sorry! I cannot your request right now!"

func (bot *Bot) registerHandlers() {
	bot.addHandler(webhookListener)
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

func webhookListener(session *discordgo.Session, message *discordgo.MessageCreate) {
	chanID := os.Getenv("PRIVATE_CHANNEL_ID")
	pubChanID := os.Getenv("PUBLIC_CHANNEL_ID")

	if message.ChannelID == chanID {

		// targetURL := ""

		// for _, emb := range message.Embeds {
		// 	if emb.URL != "" {
		// 		targetURL = emb.URL
		// 		continue
		// 	}
		// }

		// m, err := session.ChannelMessageSend(pubChanID, "Serverus received a webhook notification, this is the url "+targetURL)
		// if err != nil {
		// 	log.Println("Error sending message")
		// }

		file, err := os.Open("./images/test.png")
		if err != nil {
			log.Println("could not read image file")
		}

		message, err := session.ChannelFileSend(chanID, "image.png", file)
		if err != nil {
			panic(err)
		}

		imageURL := ""
		proxyURL := ""
		for _, image := range message.Attachments {
			imageURL = image.URL
			proxyURL = image.ProxyURL
			continue
		}

		session.ChannelMessageSendEmbed(pubChanID, &discordgo.MessageEmbed{Type: discordgo.EmbedTypeImage, Image: &discordgo.MessageEmbedImage{
			URL:      imageURL,
			ProxyURL: proxyURL,
		}})
	}
}

// type ChannelInfo struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// func helloWorld(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.Content == "Hello" {
// 		session.ChannelMessageSend(message.ChannelID, "World!")
// 	}
// }

// func channelInfo(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.Content == "Channel Info" {
// 		channelInfo, err := session.Channel(message.ChannelID)
// 		if err != nil {
// 			session.ChannelMessageSend(message.ChannelID, "Could not get channel information")
// 			return
// 		}

// 		jsonData, _ := json.MarshalIndent(ChannelInfo{
// 			ID:   channelInfo.ID,
// 			Name: channelInfo.Name,
// 		}, "", "\t")

// 		session.ChannelMessageSend(message.ChannelID, string(jsonData))
// 	}
// }
