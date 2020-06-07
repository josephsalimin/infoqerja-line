package service

import (
	"infoqerja-line/app/constant"
	iql "infoqerja-line/app/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingHelp) Reply(bot iql.BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.HelpMessage)).Do()
	return err
}
