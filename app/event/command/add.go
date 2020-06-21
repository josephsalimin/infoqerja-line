package command

import (
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingAdd : A class to represent the add job command
type IncomingAdd struct{}

// GetReply : Method service for IncomingAdd instance
func (handler *IncomingAdd) GetReply() []linebot.SendingMessage {
	// do something here
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.AddMessage), linebot.NewTextMessage("Please add job title")}
}

