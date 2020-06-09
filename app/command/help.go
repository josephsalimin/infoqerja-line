package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// GetMessage : Method service for IncomingHelp instance
func (handler *IncomingHelp) GetMessage() []linebot.SendingMessage {
	template := linebot.NewButtonsTemplate(
		"", "Help Menu", "Please click button below to refer to available command",
		linebot.NewMessageAction("View Command", "!help"),
		linebot.NewMessageAction("Add Job", "!add"),
		linebot.NewMessageAction("Show Job", "!show"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version", template)}
}
