package service

import (
	"infoqerja-line/app/constant"
	iql "infoqerja-line/app/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingUnknown :  Instance for handling unknown message
type IncomingUnknown struct{}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingUnknown) Reply(bot iql.BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.UnknownMessage)).Do()
	return err
}
