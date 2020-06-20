package input

import (
	crud "infoqerja-line/app/crud"
	model "infoqerja-line/app/model"
	constant "infoqerja-line/app/utils/constant"
	"log"
)

// IncomingStartInput : A struct to represent incoming adding date to certain job to the database by certain user
type IncomingStartInput struct {
	Data BaseData
}

// Execute : A method for Executing Starting Point job
func (job *IncomingStartInput) Execute() error {
	// get user data
	user, err := crud.ReadSingleUserData(job.Data.SourceID)

	// means that the user already applied to the
	if err == nil {

		if err = crud.DeleteJob(user.SourceID); err != nil {
			log.Print(err)
			return err
		}

		user.State = constant.WaitTitleInput
		if err = crud.UpdateUser(user); err != nil {
			log.Print(err)
			return err
		}
	}

	// creating new user data recipient
	user = model.NewUserData(job.Data.SourceID, constant.WaitTitleInput)
	if err = crud.CreateUserData(user); err != nil {
		log.Print(err)
		return err
	}

	// creating new job
	jobListing := model.NewJob("", "", "", false, job.Data.SourceID)
	if err = crud.CreateJob(jobListing); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
