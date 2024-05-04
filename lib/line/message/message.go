package message

import (
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func Push(message string, options ...PushOption) error {
	config := buildConfig(options)

	bot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return err
	}

	bot.PushMessage(&messaging_api.PushMessageRequest{
		To: os.Getenv("LINE_MESSAGE_RECEIVER"),
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Sender: config.Sender,
				Text:   message,
			},
		},
	}, "")

	return nil
}

type PushConfig struct {
	Sender *messaging_api.Sender
}

func buildConfig(options []PushOption) *PushConfig {
	config := &PushConfig{}
	for _, option := range options {
		option(config)
	}
	return config
}

type PushOption func(*PushConfig)

func WithSender(sender *messaging_api.Sender) PushOption {
	return func(c *PushConfig) {
		c.Sender = sender
	}
}
