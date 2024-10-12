package bot

import (
	"bytes"
	"log"
	"os"
	"time"

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

// This function will scrape the corresponding
// WarCraftLogs Url and generate a screenshot of the
// damage report to later embed in the specified
// channel.
func (b *Bot) generateWowReport(urlStr string) {
	// TODO - Add these channel as part of the bot configuration
	privChanID := os.Getenv("PRIVATE_CHANNEL_ID")
	pubChanID := os.Getenv("PUBLIC_CHANNEL_ID")

	if urlStr == "" {
		return
	}

	log.Println("wowReport - sleeping to await page load")
	time.Sleep(5 * time.Minute)

	// Capture the screenshot
	screenshot, reportTitle, err := captureScreenshot(urlStr)
	if err != nil {
		log.Printf("bot - error capturing screenshot for target %s , %s\n", urlStr, err.Error())
		return
	}

	// Upload the screenshot
	log.Println("wowReport - uploading screenshot to discord backup")
	imageBackupMsg, err := b.session.ChannelFileSend(privChanID, reportTitle+".png", bytes.NewReader(screenshot))
	if err != nil {
		log.Printf("bot - error saving backup of screenshot %s - %s \n", urlStr, err.Error())
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
		URL:   urlStr,
		Title: reportTitle,
		Type:  discordgo.EmbedTypeImage,
		Image: &discordgo.MessageEmbedImage{
			URL:      imageURL,
			ProxyURL: imageProxyURL,
		},
	}

	// Publish the message
	log.Println("wowReport - publishing the message")
	b.session.ChannelMessageSendEmbed(pubChanID, outMsg)

	log.Println("wowReport - published")
}
