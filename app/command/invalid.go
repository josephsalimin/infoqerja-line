package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingInvalid :  Instance for handling invalid message
type IncomingInvalid struct{}

// GetReply : Method service for IncomingInvalid instance
func (handler *IncomingInvalid) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		"", "Invalid Command", "Please click button below to refer to available command",
		linebot.NewMessageAction("View Command", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version !!", template)}
}
