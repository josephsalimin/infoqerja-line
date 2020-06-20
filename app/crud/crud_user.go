package crud

import (
	model "infoqerja-line/app/model"
	"log"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
)

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
		log.Print("debugging")
		return nil, err
	}

	return userData, nil

}

// ReadUserData : Function to use for reading all user data in the database
func ReadUserData() ([]model.UserData, error) {
	usersData := []model.UserData{}
	if err := mgm.Coll(&model.UserData{}).SimpleFind(&usersData, bson.M{}); err != nil {
		log.Print(err)
		return nil, err
	}
	return usersData, nil
}

// UpdateUser : Updating user data , usually for changing the user state when adding job
func UpdateUser(user *model.UserData) error {
	if err := mgm.Coll(user).Update(user); err != nil {
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
