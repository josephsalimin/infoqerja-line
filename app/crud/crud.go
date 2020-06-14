package crud

import (
	model "infoqerja-line/app/model"
	"log"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateJob : Creating a new job in the database, usually for starting the inserting job data process
func CreateJob(job *model.Job) error {
	jobColl := mgm.Coll(job)
	if err := jobColl.Create(job); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// CreateUserData : Creating a new user data, usually to save user state when the user inserting job using bot line
func CreateUserData(user *model.UserData) error {
	userColl := mgm.Coll(user)
	if err := userColl.Create(user); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// ReadSingleUserData : Reading the state for the user. Useful when trying to get the state the user currently in.
func ReadSingleUserData(sourceID string) (*model.UserData, error) {
	userData := &model.UserData{}

	if err := mgm.Coll(userData).First(bson.M{"sourceID": sourceID}, userData); err != nil {
		log.Print(err)
		return nil, err
	}

	return userData, nil

}

// ReadJob : Function to use for reading all completed inserted data
func ReadJob() ([]model.Job, error) {
	result := []model.Job{}
	if err := mgm.Coll(&model.Job{}).SimpleFind(&result, bson.M{
		"isComplete": true,
		"deadline": bson.M{
			"$gt": time.Now(),
		}}); err != nil {
		log.Print(err)
		return nil, err
	} // check on how to get datetime
	return result, nil
}

// ReadCurrentNotFinishedJob : Read the current job that user are trying to insert
func ReadCurrentNotFinishedJob(sourceID string) (*model.Job, error) {
	jobData := &model.Job{}
	if err := mgm.Coll(&model.Job{}).First(bson.M{
		"isComplete": false,
		"sourceID":   sourceID,
	}, jobData); err != nil {
		log.Print(err)
		return nil, err
	}

	return jobData, nil

}

// UpdateJob : Updating job , usually for completing data input from user , change is done from other methods
func UpdateJob(job *model.Job) error {
	if err := mgm.Coll(job).Update(job); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// UpdateUser : Updating user data , usually for changing the user state when adding job
func UpdateUser(user *model.UserData) error {
	if err := mgm.Coll(user).Update(user); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// DeleteJob : Deleting current job, usually because either a failure in completing data input from user, or cancelation from user when inputing data
func DeleteJob(sourceID string) error {
	jobData, err := ReadCurrentNotFinishedJob(sourceID)
	if err != nil {
		return errors.Wrap(err, "cannot find not finished job from this user")
	}

	if err = mgm.Coll(&model.Job{}).Delete(jobData); err != nil {
		log.Print(err)
		return err
	}

	return nil

}

// DeleteUserData : Deleting current job, usually because insertion has been done (successful) or the insertion is cancelled prematurely
func DeleteUserData(sourceID string) error {
	userData := &model.UserData{}
	if err := mgm.Coll(&model.UserData{}).First(bson.M{
		"sourceID": sourceID,
	}, userData); err != nil {
		log.Print(err)
		return err
	}

	if err := mgm.Coll(&model.UserData{}).Delete(userData); err != nil {
		log.Print(err)
		return err
	}

	return nil

}
