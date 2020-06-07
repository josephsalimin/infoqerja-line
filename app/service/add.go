package service

import (
	"infoqerja-line/app/constant"
	iql "infoqerja-line/app/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingAdd : A class to represent the add job command
type IncomingAdd struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingAdd) Reply(bot iql.BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.AddMessage)).Do()
	return err
}
