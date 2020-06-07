package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingShow : A class to represent the show job command
type IncomingShow struct{}

// GetMessage : Method service for IncomingHelp instance
func (handler *IncomingShow) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.ShowMessage)}
}
