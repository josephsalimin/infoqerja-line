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
		linebot.NewDatetimePickerAction("Add Date !!", "DATE", "date", "", "", ""),
		linebot.NewMessageAction("Say message", "Rice=米"),
	)

	_, err := bot.ReplyMessage(token, linebot.NewTemplateMessage("Date time picker alt text", template)).Do()
	return err
}
