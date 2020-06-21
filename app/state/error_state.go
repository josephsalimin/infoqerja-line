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

// Execute : A method for Executing Incoming Error Event job, normally for fallback from other things
func (state *ErrorState) Execute() error {
	var err error
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	// if no user detected, please delete all job created by this
	jobs, _ := (&util.JobReader{}).ReadFiltered(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	for _, job := range jobs {
		if err = job.Delete(); err != nil {
			log.Print(err)
		}
	}

	if err = user.Delete(); err != nil {
		log.Print(err)
	}

	return nil
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
	// if no user detected, please delete all job created by this
	jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	for _, job := range jobs {
		if err = job.Delete(); err != nil {
			log.Print(err)
		}
	}

	return nil

}

// NextState : Proceed to the next state
func (state *ErrorState) NextState() error {
	user := state.Data.User
	if err := user.Delete(); err != nil {
		log.Print(err)
	}
	return nil
}
