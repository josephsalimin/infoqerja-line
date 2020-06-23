package event

import (
	"infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Detail : An event representating a request of certain joblisting detail
type Detail struct {
	ObjectID primitive.ObjectID
}

// GetReply : Method service for Detail Event
func (handler *Detail) GetReply() []linebot.SendingMessage {

	template := linebot.NewButtonsTemplate(
		constant.AlertImageURL, "Invalid Command", constant.HelpShortMessage,
		linebot.NewMessageAction("Click Me", "!help"),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.InvalidMessage, template)}
}

// Parse : Method for parsing data needed for current event
func (handler *Detail) Parse(event linebot.Event) error {
	return nil
}

// Process : Method for processing current event
func (handler *Detail) Process() error {
	return nil
}
