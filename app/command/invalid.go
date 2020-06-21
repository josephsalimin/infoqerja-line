package command

import (
	model "infoqerja-line/app/model"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingInvalid :  Instance for handling invalid message
type Invalid struct{}

// GetReply : Method service for IncomingInvalid instance
func (handler *Invalid) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		"", "Invalid Command", "Please click button below to refer to available command",
		linebot.NewMessageAction("View Command", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view this in Mobile Version !!", template)}
}

func (handler *Invalid) GetState() (model.State, error) {
	return nil, nil
}
