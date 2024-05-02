package bot

func (bot *Bot) HealthCheckMessage(channel string) error {

	_, err := bot.session.ChannelMessageSend(channel, "Server is healthy")
	if err != nil {
		return err
	}

	return nil
}

func (bot *Bot) SendChannelMessage(channel string, message string) (*string, error) {

	msg, err := bot.session.ChannelMessageSend(channel, message)
	if err != nil {
		return nil, err
	}

	return &msg.ID, nil
}
