package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// GetMessage : Method service for IncomingHelp instance
func (handler *IncomingHelp) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.HelpMessage)}
}
