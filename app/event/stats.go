package event

import (
	model "infoqerja-line/app/model"
	"infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Stats : An event representating a request of statistic of current user
type Stats struct {
	Data model.BaseData
}

// GetReply : Method service for Stats event
func (handler *Stats) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		constant.AlertImageURL, "Invalid Command", constant.HelpShortMessage,
		linebot.NewMessageAction("Click Me", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.InvalidMessage, template)}
}

// Parse : Method for parsing data needed for current event
func (handler *Stats) Parse(event linebot.Event) error {
	return nil
}

// Process : Method for processing current event
func (handler *Stats) Process() error {
	return nil
}
