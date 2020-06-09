package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// GetMessage : Method service for IncomingHelp instance
func (handler *IncomingHelp) GetMessage() []linebot.SendingMessage {
	url := "https://assets.pinterest.com/ext/embed.html?id=537406168031099917"
	template := linebot.NewCarouselTemplate(
		linebot.NewCarouselColumn(
			url, "Help", "HELP",
			linebot.NewMessageAction("View Command", "!help"),
		),
		linebot.NewCarouselColumn(
			url, "Add", "ADD JOB",
			linebot.NewMessageAction("Add Job", "!add"),
		),
		linebot.NewCarouselColumn(
			url, "View", "VIEW JOB",
			linebot.NewMessageAction("Show Job", "!show"),
		),
	)

	// template1 := linebot.NewButtonsTemplate(
	// 	"", "Help Menu", "Please click button below to refer to available command",
	// 	linebot.NewMessageAction("View Command", "!help"),
	// 	linebot.NewMessageAction("Add Job", "!add"),
	// 	linebot.NewMessageAction("Show Job", "!show"),
	// )

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version", template)}
}
