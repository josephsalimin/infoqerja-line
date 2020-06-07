package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingInvalid :  Instance for handling invalid message
type IncomingInvalid struct{}

// GetMessage : Method service for IncomingInvalid instance
func (handler *IncomingInvalid) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.InvalidMessage)}
}
