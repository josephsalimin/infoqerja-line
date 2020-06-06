package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingInvalid :  Instance for handling invalid message
type IncomingInvalid struct{}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingInvalid) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(GetInvalidReplyMessage())).Do()
	return err
}

// GetInvalidReplyMessage : A function to get invalid reply message
func GetInvalidReplyMessage() string {
	message, err := GetMessageFromFile("message/invalid.txt")

	if err != nil {
		log.Print(err)
	}

	return message
}
