package command

import (
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Add : A class to represent the add job command
type Add struct{}

// GetReply : Method service for IncomingAdd instance
func (handler *Add) GetReply() []linebot.SendingMessage {
	// do something here
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.AddMessage), linebot.NewTextMessage("Please add job title")}
}

// func (handler *Add) Do() error {

// }
