package state

import (
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// IncomingAddTitleJob : A struct to represent incoming adding title to certain job to the database by certain user
type IncomingAddTitleJob struct {
	Data BaseData
}

// Execute : A method for Executing Incoming Add Title job
func (job *IncomingAddTitleJob) Execute() error {
	user, err := (&util.UserDataReader{}).ReadOne(bson.M{
		constant.SourceID: job.Data.SourceID,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	jobListing, err := (&util.JobReader{}).ReadOne(bson.M{
		constant.SourceID:   job.Data.SourceID,
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
	jobListing.Title = job.Data.Input
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
