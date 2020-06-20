package input

import (
	crud "infoqerja-line/app/crud"
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"
)

// IncomingAddDateJob : A struct to represent incoming adding date to certain job to the database by certain user
type IncomingAddDateJob struct {
	Data BaseData
}

func (job *IncomingAddDateJob) Execute() error {
	user, err := crud.ReadSingleUserData(job.Data.SourceID)
	jobListing, err := crud.ReadCurrentNotFinishedJob(job.Data.SourceID)

	if err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	t, err := time.Parse(constant.DateFormatLayout, job.Data.Input)
	if err != nil {
		t = time.Now()
		log.Print(err)
	}

	jobListing.Deadline = t
	jobListing.IsComplete = true
	if err = crud.UpdateJob(jobListing); err != nil {
		log.Print(err)
		return err
	}

	if err = crud.DeleteUserData(user.SourceID); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
