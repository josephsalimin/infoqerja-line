package event

import (
	"infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	"infoqerja-line/app/utils/constant"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Detail : An event representating a request of certain joblisting detail
type Detail struct {
	objectID primitive.ObjectID
	job      model.Job
}

// GetReply : Method service for Detail Event
func (handler *Detail) GetReply() []linebot.SendingMessage {

	// template := linebot.NewButtonsTemplate(
	// 	constant.AlertImageURL, "Invalid Command", constant.HelpShortMessage,
	// 	linebot.NewMessageAction("Click Me", "!help"),
	// )

	contents := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Hello,",
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "World!",
				},
			},
		},
	}
	return []linebot.SendingMessage{linebot.NewFlexMessage(constant.InvalidMessage, contents)}
}

// Parse : Method for parsing data needed for current event
func (handler *Detail) Parse(event linebot.Event) error {
	id, err := primitive.ObjectIDFromHex(strings.Split(event.Postback.Data, "|")[1])
	if err != nil {
		log.Print(err)
	} else {
		handler.objectID = id
	}
	return err

}

// Process : Method for processing current event
func (handler *Detail) Process() error {
	job, err := (&util.JobReader{}).ReadOne(bson.M{
		"_id": handler.objectID,
	})
	if err != nil {
		log.Print(err)
	} else {
		handler.job = *job
	}
	return err
}
