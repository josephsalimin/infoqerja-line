package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingHelp) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(GetHelpReplyMessage())).Do()
	return err
}

// GetHelpReplyMessage : A function to get help reply message
func GetHelpReplyMessage() string {
	message, err := GetMessageFromFile("./message/help.txt")

	if err != nil {
		log.Print(err)
	}

	return message
}
