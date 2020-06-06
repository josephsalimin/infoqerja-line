package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingUnknown :  Instance for handling unknown message
type IncomingUnknown struct{}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingUnknown) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(GetUnknownReplyMessage())).Do()
	return err
}

// GetUnknownReplyMessage : A function to get unknown reply message
func GetUnknownReplyMessage() string {
	message, err := GetMessageFromFile("./message/unknown.txt")

	if err != nil {
		log.Print(err)
	}

	return message
}
