package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingAdd : A class to represent the add job command
type IncomingAdd struct{}

// GetReply : Method service for IncomingHelp instance
func (handler *IncomingAdd) GetReply() []linebot.SendingMessage {
	template := linebot.NewButtonsTemplate(
		"", "Adding Job", "Select job date deadline!",
		linebot.NewDatetimePickerAction("Add Date", "DATE", "date", "", "", ""),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version !!", template)}
}
