package line

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingUnknown :  Instance for handling unknown message
type IncomingUnknown struct{}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingUnknown) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.UnknownMessage)).Do()
	return err
}
