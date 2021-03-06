package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// AddDescState : A struct to represent incoming adding description to certain job to the database by certain user
type AddDescState struct {
	Data model.BaseData
}

// GetReply : Get the reply for next question
func (state *AddDescState) GetReply() []linebot.SendingMessage {
	template := linebot.NewButtonsTemplate(
		constant.ClockImageURL, "", constant.AddDateMessage,
		linebot.NewDatetimePickerAction("Pick a date", constant.DateData, "date", "", "", ""),
	)

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.UnavailableMessage, template)}
}

// Parse : Parse data needed by certain state
func (state *AddDescState) Parse(event linebot.Event) error {
	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
		Input:    util.GetData(event.Message),
	}

	return nil
}

// Process : Do certain process for certain state
func (state *AddDescState) Process() error {

	jobListing, err := (&util.JobReader{}).ReadOne(bson.M{
		constant.SourceID:   state.Data.SourceID,
		constant.IsComplete: false,
	})

	if err != nil {
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

// NextState : Proceed to the next state
func (state *AddDescState) NextState() error {
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	user.State = constant.WaitDateInput
	if err := user.Update(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
