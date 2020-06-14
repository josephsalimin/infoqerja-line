package model

import (
	"github.com/bykof/stateful"
)

// UserData : A model to represent the user data in the database
type UserData struct {
	State    stateful.State
	SourceId string
}

// NewUserData : default constructor for UserData struct
func NewUserData(source string, state stateful.State) *UserData {
	return &UserData{
		SourceId: source,
		State:    state,
	}
}
