package model

import "github.com/line/line-bot-sdk-go/linebot"

// BaseData : the basec structure for data (DTO) in the event layer
type BaseData struct {
	SourceID string
	Input    string
	User     UserData
}

type (
	// Processor : Simple interface for doing certain process
	Processor interface {
		Process() error
	}
	// Stater
	Stater interface {
		GetState() (State, error)
	}
	// Replier : Simple interface for getting reply context
	Replier interface {
		GetReply() []linebot.SendingMessage
	}
	// Parser : Simple interface for parsing data
	Parser interface {
		Parse(event linebot.Event) error
	}
	// Next : Simple interface to move into next state (state design pattern)
	Next interface {
		NextState() error
	}
)

// Command : An interface for command struct that able to do certain action and have reply
type Command interface {
	Stater
	Replier
}

// State : An interface for certain state when user adding job that were able to reply, parse, do certain action, and proceed to next state
type State interface {
	Replier
	Parser
	Processor
	Next
}
