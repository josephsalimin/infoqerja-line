package service

import (
	iql "infoqerja-line/app/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingAdd : A class to represent the add job command
type IncomingAdd struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingAdd) Reply(bot iql.BotClient, token string) error {

	template := linebot.NewButtonsTemplate(
		"", "Adding Job!!", "Select job date deadline!",
		linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
	)

	_, err := bot.ReplyMessage(token, linebot.NewTemplateMessage("Date time picker alt text", template)).Do()
	return err
}
