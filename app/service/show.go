package service

import (
	"infoqerja-line/app/constant"
	iql "infoqerja-line/app/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingShow : A class to represent the show job command
type IncomingShow struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingShow) Reply(bot iql.BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.ShowMessage)).Do()
	return err
}
