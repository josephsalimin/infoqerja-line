package model

import "github.com/Kamva/mgm/v2"

// UserData : A model to represent the user data in the database
type UserData struct {
	mgm.DefaultModel `bson:",inline"`
	State            string `json:"state" bson:"state"`
	SourceID         string `json:"sourceID" bson:"sourceID"`
}

// NewUserData : default constructor for UserData struct
func NewUserData(source string, state string) *UserData {
	return &UserData{
		SourceID: source,
		State:    state,
	}
}

// GetState implement interface stateful
func (user UserData) GetState() string {
	return user.State
}

// SetState implement interface stateful
func (user *UserData) SetState(state string) error {
	user.State = state
	return nil
}
