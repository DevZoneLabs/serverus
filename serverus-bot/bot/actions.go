package bot

func (bot *Bot) HealthCheckMessage(channel string) error {
	
	_, err := bot.session.ChannelMessageSend(channel, "Server is healthy")
	if err != nil {
		return err
	}

	return nil
}