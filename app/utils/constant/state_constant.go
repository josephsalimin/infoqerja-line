package constant

import "github.com/bykof/stateful"

const (
	// WaitTitleInput : state where the machine / bot expect a source to input the job title
	WaitTitleInput = stateful.DefaultState("WaitTitleInput")

	// WaitDescInput : a state where the machine / bot expect a source to input the description of the job
	WaitDescInput = stateful.DefaultState("WaitDescInput")

	// WaitDateInput : a state where the machine / bot expect a source to input the datetime for the job
	WaitDateInput = stateful.DefaultState("WaitDateInput")
)
