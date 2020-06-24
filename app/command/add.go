package command

import (
	model "infoqerja-line/app/model"
	"infoqerja-line/app/state"
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Add : A class to represent the add job command
type Add struct{}

// GetReply : Method service for IncomingAdd instance
func (handler *Add) GetReply() []linebot.SendingMessage {
	// do something here
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.AddMessage)}
}

// GetState : Method to get any state a certain command produce, if present
func (handler *Add) GetState() (model.State, error) {
	return &state.StartState{}, nil
}
