package model

import (
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"

	"github.com/Kamva/mgm/v2"
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
		log.Print(err)
	}

	return &Job{
		SourceID:    sourceID,
		Title:       title,
		Deadline:    t,
		IsComplete:  check,
		Description: desc,
	}
}
