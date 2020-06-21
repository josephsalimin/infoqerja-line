package command

import (
	model "infoqerja-line/app/model"
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Welcome : A class to represent the welcoming user event
type Welcome struct{}

// GetReply : Method service for Welcome instance
func (handler *Welcome) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.WelcomeMessage)}
}

func (handler *Welcome) GetState() (model.State, error) {
	return nil, nil
}

// UnWelcome : A class to represent the part away user event
type UnWelcome struct{}

// GetReply : Method service for UnWelcome instance
func (handler *UnWelcome) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.UnWelcomeMessage)}
}

func (handler *UnWelcome) GetState() (model.State, error) {
	return nil, nil
}
