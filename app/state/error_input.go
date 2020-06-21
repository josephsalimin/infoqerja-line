package state

import (
	crud "infoqerja-line/app/crud"
	"log"
)

// IncomingErrorEvent : A struct to represent incoming error event to certain job to the database by certain user
type IncomingErrorEvent struct {
	Data BaseData
}

// Execute : A method for Executing Incoming Error Event job
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
