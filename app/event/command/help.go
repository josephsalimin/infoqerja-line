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
			"", "HELP", "Press the button below to show InfoQerja Menu",
			linebot.NewMessageAction("Click Me!", "!help"),
		),
		linebot.NewCarouselColumn(
			"", "ADD JOB", "Press the button below to add a job",
			linebot.NewMessageAction("Click Me!", "!add"),
		),
		linebot.NewCarouselColumn(
			"", "VIEW JOB", "Press the button below to show joblist",
			linebot.NewMessageAction("Click Me!", "!show"),
		),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version", template)}
}
