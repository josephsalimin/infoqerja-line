package state

import (
	crud "infoqerja-line/app/crud"
	constant "infoqerja-line/app/utils/constant"
	"log"
)

// IncomingAddTitleJob : A struct to represent incoming adding title to certain job to the database by certain user
type IncomingAddTitleJob struct {
	Data BaseData
}

// Execute : A method for Executing Incoming Add Title job
func (job *IncomingAddTitleJob) Execute() error {
	user, err := crud.ReadSingleUserData(job.Data.SourceID)

	if err != nil {
		log.Print(err)
		return err
	}

	jobListing, err := crud.ReadCurrentNotFinishedJob(job.Data.SourceID)

	if err != nil {
		log.Print(err)
		return err
	}

	// update current state
	user.State = constant.WaitDescInput
	if err = crud.UpdateUser(user); err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Title = job.Data.Input
	if err = crud.UpdateJob(jobListing); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
