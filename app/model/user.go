package model

// UserData : A model to represent the user data in the database
type UserData struct {
	state    string
	SourceID string
}

// NewUserData : default constructor for UserData struct
func NewUserData(source string, state string) *UserData {
	return &UserData{
		SourceID: source,
		state:    state,
	}
}

// GetState implement interface stateful
func (user UserData) GetState() string {
	return user.state
}

// SetState implement interface stateful
func (user *UserData) SetState(state string) error {
	user.state = state
	return nil
}
