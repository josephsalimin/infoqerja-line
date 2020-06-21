package state

import (
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// IncomingAddDateJob : A struct to represent incoming adding date to certain job to the database by certain user
type IncomingAddDateJob struct {
	Data BaseData
}

// Execute : A method for Executing Incoming Add Date job
func (job *IncomingAddDateJob) Execute() error {
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

	// update joblisting data
	t, err := time.Parse(constant.DateFormatLayout, job.Data.Input)
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

	if err = user.Delete(); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
