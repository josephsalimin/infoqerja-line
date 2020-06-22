package command

import (
	model "infoqerja-line/app/model"
	"infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Invalid :  Instance for handling invalid message
type Invalid struct{}

// GetReply : Method service for IncomingInvalid instance
func (handler *Invalid) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		"https://img.icons8.com/material-outlined/96/000000/warning-shield.png", "Invalid Command", constant.HelpShortMessage,
		linebot.NewMessageAction("Click Me", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.InvalidMessage, template)}
}

// GetState : Method to get any state a certain command produce, if present
func (handler *Invalid) GetState() (model.State, error) {
	return nil, nil
}
