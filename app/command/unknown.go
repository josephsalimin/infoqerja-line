package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingUnknown :  Instance for handling unknown message
type IncomingUnknown struct{}

// GetMessage : Method service for IncomingInvalid instance
func (handler *IncomingUnknown) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.UnknownMessage)}
}
