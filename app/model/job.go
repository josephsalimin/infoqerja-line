package model

import (
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"
)

// Job : A model to represent the job data in the database
type Job struct {
	SourceID    string
	Deadline    time.Time
	Description string
	Title       string
	isFinish    bool
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
		isFinish:    check,
		Description: desc,
	}
}
