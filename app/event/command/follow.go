package command

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Welcome : A class to represent the welcoming user event
type Welcome struct{}

// GetMessage : Method service for Welcome instance
func (handler *Welcome) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.WelcomeMessage)}
}

// UnWelcome : A class to represent the part away user event
type UnWelcome struct{}

// GetMessage : Method service for UnWelcome instance
func (handler *UnWelcome) GetMessage() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.UnWelcomeMessage)}
}
