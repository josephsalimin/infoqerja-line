package model

import (
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"

	"github.com/Kamva/mgm/v3"
)

// Job : A model to represent the job data in the database
type Job struct {
	mgm.DefaultModel `bson:",inline"`
	SourceID         string    `json:"sourceID" bson:"sourceID"`
	Deadline         time.Time `json:"deadline" bson:"deadline"`
	Description      string    `json:"desc" bson:"desc"`
	Title            string    `json:"title" bson:"title"`
	IsComplete       bool      `json:"isComplete" bson:"isComplete"`
}

// NewJob : default constructor for Job struct
func NewJob(date string, desc string, title string, check bool, sourceID string) *Job {
	t, err := time.Parse(constant.DateFormatLayout, date)
	if err != nil {
		t = time.Now()
	}
	return &Job{
		SourceID:    sourceID,
		Title:       title,
		Deadline:    t,
		IsComplete:  check,
		Description: desc,
	}
}

// Create : An method to insert the data into database
func (job *Job) Create() error {
	if err := mgm.Coll(job).Create(job); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// Update : An method to update the data into database
func (job *Job) Update() error {
	if err := mgm.Coll(job).Update(job); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// Delete : An method to delete the data in the database
func (job *Job) Delete() error {
	if err := mgm.Coll(&Job{}).Delete(job); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
