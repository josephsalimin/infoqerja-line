package input

import (
	crud "infoqerja-line/app/crud"
	constant "infoqerja-line/app/utils/constant"
	"log"
)

type IncomingAddDescJob struct {
	Data BaseData
}

func (job *IncomingAddDescJob) Execute() error {
	user, err := crud.ReadSingleUserData(job.Data.SourceID)
	jobListing, err := crud.ReadCurrentNotFinishedJob(job.Data.SourceID)

	if err != nil {
		log.Print(err)
		return err
	}

	// update current state
	user.State = constant.WaitDateInput
	if err = crud.UpdateUser(user); err != nil {
		log.Print(err)
		return err
	}

	// update joblisting data
	jobListing.Description = job.Data.Input
	if err = crud.UpdateJob(jobListing); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
