package input

// StartInput :
type IncomingStartInput struct{}

func (job *IncomingStartInput) Execute() error {
	return nil
}
