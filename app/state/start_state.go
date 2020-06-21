package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// IncomingStartInput : A struct to represent incoming adding date to certain job to the database by certain user
type IncomingStartInput struct {
	Data model.BaseData
}

// Execute : A method for Executing Starting Point job
func (job *IncomingStartInput) Execute() error {

	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: job.Data.SourceID,
	})

	if err != nil { // if no user data detected
		if err = model.NewUserData(job.Data.SourceID, constant.WaitTitleInput).Create(); err != nil {
			log.Print(err)
			return err
		}
	} else { // if user data detected, then please update the data, and delete previous inputed job data that were incomplete
		jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
			constant.SourceID:   job.Data.SourceID,
			constant.IsComplete: false,
		})

		for _, job := range jobs {
			if err = job.Delete(); err != nil {
				log.Print(err)
			}
		}

		user.State = constant.WaitTitleInput
		if err = user.Update(); err != nil {
			log.Print(err)
			return err
		}

	}

	if err = model.NewJob("", "", "", false, job.Data.SourceID).Create(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
