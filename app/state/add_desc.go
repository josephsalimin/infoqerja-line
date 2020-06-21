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

// AddDescState : A struct to represent incoming adding description to certain job to the database by certain user
type AddDescState struct {
	Data model.BaseData
}

// Execute : A method for Executing Incoming Add Desc job
func (state *AddDescState) Execute() error {
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
	user.State = constant.WaitDateInput
	if err = user.Update(); err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Description = state.Data.Input
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetReply : Get the reply for next question
func (state *AddDescState) GetReply() []linebot.SendingMessage {
	template := linebot.NewButtonsTemplate(
		"", "", "Select date for job deadline :)",
		linebot.NewDatetimePickerAction("Pick Date", "DATE", "date", time.Now().String(), "", ""),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage("Please view in mobile device", template)}
}

// Parse : Parse data needed by certain state
func (state *AddDescState) Parse(event linebot.Event) error {
	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: util.GetSource(event),
	})
	if err != nil {
		log.Print(err)
		return err
	}

	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
		Input:    util.GetData(event.Message),
		User:     *user,
	}

	return nil
}

// Process : Do certain process for certain state
func (state *AddDescState) Process() error {
	_, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	if err != nil { // if no user data detected
		if err = model.NewUserData(state.Data.SourceID, constant.WaitTitleInput).Create(); err != nil {
			log.Print(err)
			return err
		}
	} else { // if user data detected, then please update the data, and delete previous inputed state data that were incomplete
		jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
			constant.SourceID:   state.Data.SourceID,
			constant.IsComplete: false,
		})

		for _, state := range jobs {
			if err = state.Delete(); err != nil {
				log.Print(err)
			}
		}
	}

	// Creating new job
	if err = model.NewJob("", "", "", false, state.Data.SourceID).Create(); err != nil {
		log.Print(err)
		return err
	}

	return nil

}

// NextState : Proceed to the next state
func (state *AddDescState) NextState() error {
	user := state.Data.User
	user.State = constant.WaitDateInput
	if err := user.Update(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
