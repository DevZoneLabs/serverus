package bot

import (
	"bytes"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) HealthCheckMessage(channel string) error {
	_, err := b.session.ChannelMessageSend(channel, "Server is healthy")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) SendChannelMessage(channel string, message string) (*string, error) {
	msg, err := b.session.ChannelMessageSend(channel, message)
	if err != nil {
		return nil, err
	}

	return &msg.ID, nil
}

func (b *Bot) generateWowReport(msg *discordgo.MessageCreate) {
	// TODO - Add these channel as part of the bot configuration
	privChanID := os.Getenv("PRIVATE_CHANNEL_ID")
	pubChanID := os.Getenv("PUBLIC_CHANNEL_ID")

	// Check for the internal message embed. It is always going to be one
	// since the channel will be waiting for webhook calls.
	// We could modify this in the future to determine which
	// webhook comes from
	msgEmbed := &discordgo.MessageEmbed{}
	for _, emb := range msg.Embeds {
		msgEmbed = emb
	}
	if msgEmbed == nil || msgEmbed.URL == "" {
		return
	}

	// Capture the screenshot
	screenshot, reportTitle, err := captureScreenshot(msgEmbed.URL)
	if err != nil {
		log.Printf("bot - error capturing screenshot for target %s , %s\n", msgEmbed.URL, err.Error())
		return
	}

	// Upload the screenshot
	log.Println("bot - capturing screenshot for ", msgEmbed.URL)
	imageBackupMsg, err := b.session.ChannelFileSend(privChanID, "report.png", bytes.NewReader(screenshot))
	if err != nil {
		log.Printf("bot - error saving backup of screenshot %s - %s \n", msgEmbed.URL, err.Error())
		return
	}

	// Get the imagURL and ProxyURL from the backup image
	var imageURL, imageProxyURL string
	for _, attachment := range imageBackupMsg.Attachments {
		imageURL = attachment.URL
		imageProxyURL = attachment.ProxyURL
	}

	// Construct outbound message to public channel
	outMsg := &discordgo.MessageEmbed{
		URL:   msgEmbed.URL,
		Title: reportTitle,
		Type:  discordgo.EmbedTypeImage,
		Image: &discordgo.MessageEmbedImage{
			URL:      imageURL,
			ProxyURL: imageProxyURL,
		},
	}

	// Publish the message
	b.session.ChannelMessageSendEmbed(pubChanID, outMsg)
}
