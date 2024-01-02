package message

import (
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func Push(message string) error {
	bot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return err
	}

	bot.PushMessage(&messaging_api.PushMessageRequest{
		To: os.Getenv("LINE_MESSAGE_RECEIVER"),
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: message,
			},
		},
	}, "")

	return nil
}
