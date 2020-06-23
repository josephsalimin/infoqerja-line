package event

import (
	model "infoqerja-line/app/model"
	"infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Cancel : An event representating a request to cancel adding joblist in database for certain user
type Cancel struct {
	Data model.BaseData
}

// GetReply : Method service for Cancel event
func (handler *Cancel) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		constant.AlertImageURL, "Invalid Command", constant.HelpShortMessage,
		linebot.NewMessageAction("Click Me", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.InvalidMessage, template)}
}

// Parse : Method for parsing data needed for current event
func (handler *Cancel) Parse(event linebot.Event) error {
	return nil
}

// Process : Method for processing current event
func (handler *Cancel) Process() error {
	return nil
}
