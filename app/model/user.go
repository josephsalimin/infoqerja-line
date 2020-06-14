package model

import (
	"github.com/bykof/stateful"
)

// UserData : A model to represent the user data in the database
type UserData struct {
	state    stateful.State
	SourceID string
}

// NewUserData : default constructor for UserData struct
func NewUserData(source string, state stateful.State) *UserData {
	return &UserData{
		SourceID: source,
		state:    state,
	}
}

// GetState implement interface stateful
func (user UserData) GetState() stateful.State {
	return user.state
}

// SetState implement interface stateful
func (user *UserData) SetState(state stateful.State) error {
	user.state = state
	return nil
}
