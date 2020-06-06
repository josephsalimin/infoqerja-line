package line

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingInvalid :  Instance for handling invalid message
type IncomingInvalid struct{}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingInvalid) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.InvalidMessage)).Do()
	return err
}
