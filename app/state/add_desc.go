package state

import (
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// IncomingAddDescJob : A struct to represent incoming adding description to certain job to the database by certain user
type IncomingAddDescJob struct {
	Data BaseData
}

// Execute : A method for Executing Incoming Add Desc job
func (job *IncomingAddDescJob) Execute() error {
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
	user.State = constant.WaitDateInput
	if err = user.Update(); err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Description = job.Data.Input
	if err = jobListing.Update(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
