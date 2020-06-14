package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Welcome : A class to represent the welcoming user event
type Welcome struct{}

// GetReply : Method service for Welcome instance
func (handler *Welcome) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.WelcomeMessage)}
}

// UnWelcome : A class to represent the part away user event
type UnWelcome struct{}

// GetReply : Method service for UnWelcome instance
func (handler *UnWelcome) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.UnWelcomeMessage)}
}
