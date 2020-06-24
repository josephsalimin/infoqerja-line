package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// ErrorState : A struct to represent incoming error event to certain job to the database by certain user
type ErrorState struct {
	Data model.BaseData
}

// GetReply : Get the reply for next question
func (state *ErrorState) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage("An error happened. The transaction is undone due to certain process!! Please contact customer service for further info")}
}

// Parse : Parse data needed by certain state
func (state *ErrorState) Parse(event linebot.Event) error {
	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
	}
	return nil
}

// Process : Do certain process for certain state
func (state *ErrorState) Process() error {
	jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	for _, job := range jobs {
		if err = job.Delete(); err != nil {
			log.Print(err)
		}
	}

	if user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	}); err == nil && user != nil {
		user.State = constant.NoState
		if err := user.Update(); err != nil {
			log.Print(err)
		}
	}
	return nil
}

// NextState : Proceed to the next state
func (state *ErrorState) NextState() error {
	return nil
}
