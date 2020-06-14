package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingShow : A class to represent the show job command
type IncomingShow struct{}

// GetReply : Method service for IncomingHelp instance
func (handler *IncomingShow) GetReply() []linebot.SendingMessage {

	// template : carousel

	return []linebot.SendingMessage{linebot.NewTextMessage(constant.ShowMessage)}
}
