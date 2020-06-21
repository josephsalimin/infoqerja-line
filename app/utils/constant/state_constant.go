package constant

const (
	// WaitTitleInput : state where the machine / bot expect a source to input the job title
	WaitTitleInput = "wait-title"

	// WaitDescInput : a state where the machine / bot expect a source to input the description of the job
	WaitDescInput = "wait-desc"

	// WaitDateInput : a state where the machine / bot expect a source to input the datetime for the job
	WaitDateInput = "wait-date"

	// NoState : a state where user data is not registered / not accountable to any kind of transaction in database
	NoState = "non"
)
