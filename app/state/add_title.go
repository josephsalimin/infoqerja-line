package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// AddTitleState : A struct to represent incoming adding title to certain job to the database by certain user
type AddTitleState struct {
	Data model.BaseData
}

// Execute : A method for Executing Incoming Add Title job
func (state *AddTitleState) Execute() error {
	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	jobListing, err := (&util.JobReader{}).ReadOne(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	// update current state
	user.State = constant.WaitDescInput
	if err = user.Update(); err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Title = state.Data.Input
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetReply : Get the reply for next question
func (state *AddTitleState) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage("Please add job description : ")}
}

// Parse : Parse data needed by certain state
func (state *AddTitleState) Parse(event linebot.Event) error {
	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
		Input:    util.GetData(event.Message),
	}

	return nil
}

// Process : Do certain process for certain state
func (state *AddTitleState) Process() error {

	jobListing, err := (&util.JobReader{}).ReadOne(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Title = state.Data.Input
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// NextState : Proceed to the next state
func (state *AddTitleState) NextState() error {
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	user.State = constant.WaitDescInput
	if err := user.Update(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
