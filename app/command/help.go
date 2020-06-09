package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// GetMessage : Method service for IncomingHelp instance
func (handler *IncomingHelp) GetMessage() []linebot.SendingMessage {
	template := linebot.NewCarouselTemplate(
		linebot.NewCarouselColumn(
			"", "Help Command", "",
			linebot.NewMessageAction("View Command", "!help"),
		),
		linebot.NewCarouselColumn(
			"", "Add Job Command", "",
			linebot.NewMessageAction("Add Job", "!add"),
		),
		linebot.NewCarouselColumn(
			"", "View Job Command", "",
			linebot.NewMessageAction("Show Job", "!show"),
		),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version", template)}
}
