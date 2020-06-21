package crud

import (
	model "infoqerja-line/app/model"
	"log"
	"time"

	"github.com/Kamva/mgm/v3"
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
