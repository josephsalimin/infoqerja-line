package utils

import (
	model "infoqerja-line/app/model"
	"log"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	// JobReader : A struct to represent reader for Job Data
	JobReader struct{}

	// UserDataReader : A struct to represent reader for UserData
	UserDataReader struct{}
)

// ReadOne : Function to read only one data for certain model class
func (reader *JobReader) ReadOne(filter bson.M) (*model.Job, error) {
	jobData := &model.Job{}
	if err := mgm.Coll(&model.Job{}).First(filter, jobData); err != nil {
		log.Print(err)
		return nil, err
	}
	return jobData, nil
}

// ReadAll : Function to read all data in the database for certain model class
func (reader *JobReader) ReadAll() ([]model.Job, error) {
	return reader.ReadFiltered(bson.M{})
}

// ReadFiltered : Function to read all data based on certain filter for a model class
func (reader *JobReader) ReadFiltered(filter bson.M) ([]model.Job, error) {
	result := []model.Job{}
	if err := mgm.Coll(&model.Job{}).SimpleFind(&result, filter); err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

// ReadOne : Function to read only one data for certain model class
func (reader *UserDataReader) ReadOne(filter bson.M) (*model.UserData, error) {
	userData := &model.UserData{}
	if err := mgm.Coll(&model.UserData{}).First(filter, userData); err != nil {
		log.Print(err)
		return nil, err
	}
	return userData, nil
}

// ReadAll : Function to read all data in the database for certain model class
func (reader *UserDataReader) ReadAll() ([]model.UserData, error) {
	return reader.ReadFiltered(bson.M{})
}

// ReadFiltered : Function to read all data based on certain filter for a model class
func (reader *UserDataReader) ReadFiltered(filter bson.M) ([]model.UserData, error) {
	result := []model.UserData{}
	if err := mgm.Coll(&model.UserData{}).SimpleFind(&result, filter); err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}
