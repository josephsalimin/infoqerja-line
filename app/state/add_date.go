package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// AddDateState : A struct to represent incoming adding date to certain job to the database by certain user
type AddDateState struct {
	Data model.BaseData
}

// GetReply : Get the reply for next question
func (state *AddDateState) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage(constant.ThankYouMessage)}
}

// Parse : Parse data needed by certain state
func (state *AddDateState) Parse(event linebot.Event) error {
	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
		Input:    util.GetData(event.Message),
	}

	return nil
}

// Process : Do certain process for certain state
func (state *AddDateState) Process() error {
	jobListing, err := (&util.JobReader{}).ReadOne(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	t, err := time.Parse(constant.DateFormatLayout, state.Data.Input)
	if err != nil {
		t = time.Now().AddDate(0, 6, 0) // default value : deadline 6 month from now
		log.Print(err)
	}
	jobListing.Deadline = t
	jobListing.IsComplete = true
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// NextState : Proceed to the next state
func (state *AddDateState) NextState() error {
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	if err := user.Delete(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
