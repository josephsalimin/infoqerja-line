package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// StartState : A struct to represent incoming adding date to certain state to the database by certain user
type StartState struct {
	Data model.BaseData
}

// Execute : A method for Executing Starting Point state
func (state *StartState) Execute() error {

	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	if err != nil { // if no user data detected
		if err = model.NewUserData(state.Data.SourceID, constant.WaitTitleInput).Create(); err != nil {
			log.Print(err)
			return err
		}
	} else { // if user data detected, then please update the data, and delete previous inputed state data that were incomplete
		states, err := (&util.JobReader{}).ReadFiltered(bson.M{
			constant.SourceID:   state.Data.SourceID,
			constant.IsComplete: false,
		})

		for _, state := range states {
			if err = state.Delete(); err != nil {
				log.Print(err)
			}
		}

		user.State = constant.WaitTitleInput
		if err = user.Update(); err != nil {
			log.Print(err)
			return err
		}

	}

	if err = model.NewJob("", "", "", false, state.Data.SourceID).Create(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetReply : Get the reply for next question
func (state *StartState) GetReply() []linebot.SendingMessage {
	return []linebot.SendingMessage{linebot.NewTextMessage("Please add job title : ")}
}

// Parse : Parse data needed by certain state
func (state *StartState) Parse(event linebot.Event) error {
	state.Data = model.BaseData{
		SourceID: util.GetSource(event),
	}
	return nil
}

// Process : Do certain process for certain state
func (state *StartState) Process() error {
	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	if user != nil { // if no user data detected
		// if user data detected, please check the state first
		if err := user.Delete(); err != nil {
			log.Print(err)
			return err
		}
		jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
			constant.SourceID:   state.Data.SourceID,
			constant.IsComplete: false,
		})
		// be careful: bug might lurking here
		for _, job := range jobs {
			if err = job.Delete(); err != nil {
				log.Print(err)
			}
		}
	}
	user = model.NewUserData(state.Data.SourceID, constant.NoState)
	if err = user.Create(); err != nil {
		log.Print(err)
		log.Print("Creating New User")
		return err
	}
	// Creating new job
	if err = model.NewJob("", "", "", false, state.Data.SourceID).Create(); err != nil {
		log.Print(err)
		log.Print("Creating New Job")
		return err
	}

	return nil

}

// NextState : Proceed to the next state
func (state *StartState) NextState() error {
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: state.Data.SourceID,
	})

	user.State = constant.WaitTitleInput
	if err := user.Update(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
