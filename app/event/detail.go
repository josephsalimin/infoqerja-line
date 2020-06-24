package event

import (
	"infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
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

	return []linebot.SendingMessage{linebot.NewTextMessage(handler.job.Description)}
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
