package utils

import (
	iqi "infoqerja-line/app/event/input"
	constant "infoqerja-line/app/utils/constant"
)

type (
	// Job : Interface for Executing a job
	Job interface {
		Execute() error
	}

	// FinderJob : interface of searching job service
	FinderJob interface {
		GetJob() Command
	}

	// CurrState : struct representing current state of user
	CurrState struct {
		State string
	}
)

// GetState : get the type of command from user inputs
func (state *CurrState) GetState() Job {
	switch state.State {
	case constant.WaitDateInput:
		return &iqi.IncomingAddDateJob{}
	case constant.WaitDescInput:
		return &iqi.IncomingAddDescJob{}
	case constant.WaitTitleInput:
		return &iqi.IncomingAddTitleJob{}
	case constant.NoState:
		return &iqi.IncomingStartInput{}
	default:
		return &iqi.IncomingErrorInput{}
	}
}
