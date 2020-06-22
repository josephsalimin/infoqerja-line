package command

import (
	model "infoqerja-line/app/model"
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Show : A class to represent the show job command
type Show struct{}

// GetReply : Method service for IncomingHelp instance
func (handler *Show) GetReply() []linebot.SendingMessage {

	// template : carousel

	return []linebot.SendingMessage{linebot.NewTextMessage(constant.ShowMessage)}
}

// Get Data

// GetState : Method to get any state a certain command produce, if present
func (handler *Show) GetState() (model.State, error) {
	return nil, nil
}
