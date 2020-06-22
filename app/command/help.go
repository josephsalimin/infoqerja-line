package command

import (
	model "infoqerja-line/app/model"
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Help : A class to represent the help command
type Help struct{}

// GetReply : Method service for IncomingHelp instance
func (handler *Help) GetReply() []linebot.SendingMessage {
	template := linebot.NewButtonsTemplate(
		"https://img.icons8.com/ios/50/000000/about.png", "Help", "Menu",
		linebot.NewMessageAction("Help", "!help"),
		linebot.NewMessageAction("Add Job", "!add"),
		linebot.NewMessageAction("Show Job", "!show"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.HelpMessage, template)}
}

// GetState : Method to get any state a certain command produce, if present
func (handler *Help) GetState() (model.State, error) {
	return nil, nil
}
