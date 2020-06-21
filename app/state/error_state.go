package state

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// IncomingErrorEvent : A struct to represent incoming error event to certain job to the database by certain user
type IncomingErrorEvent struct {
	Data model.BaseData
}

// Execute : A method for Executing Incoming Error Event job, normally for fallback from other things
func (job *IncomingErrorEvent) Execute() error {
	var err error
	user, _ := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: job.Data.SourceID,
	})

	// if no user detected, please delete all job created by this
	jobs, _ := (&util.JobReader{}).ReadFiltered(bson.M{
		constant.SourceID:   job.Data.SourceID,
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
