package input

import (
	crud "infoqerja-line/app/crud"
	"log"
)

type IncomingErrorEvent struct {
	Data BaseData
}

func (job *IncomingErrorEvent) Execute() error {
	// get user information
	user, err := crud.ReadSingleUserData(job.Data.SourceID)
	if err != nil {
		log.Print(err)
		return err
	}

	// deleting all job data
	if err = crud.DeleteJob(user.SourceID); err != nil {
		log.Print(err)
		return err
	}

	// deleting the user data
	if err = crud.DeleteUserData(user.SourceID); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
