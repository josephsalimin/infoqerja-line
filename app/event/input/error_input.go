package input

type IncomingErrorInput struct{}

func (job *IncomingErrorInput) Execute() error {
	return nil
}
