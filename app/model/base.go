package model

import "github.com/line/line-bot-sdk-go/linebot"

// BaseData : the basec structure for data (DTO) in the event layer
type BaseData struct {
	SourceID string
	Input    string
}

type (
	// Actioner : Simple interface for doing certain action
	Actioner interface {
		Do() error
	}
	// Replier : Simple interface for getting reply context
	Replier interface {
		GetReply() []linebot.SendingMessage
	}
	// Parser : Simple interface for parsing data
	Parser interface {
		Parse() error
	}
	// Next : Simple interface to move into next state (state design pattern)
	Next interface {
		NextState() error
	}
)

// Command : An interface for command struct that able to do certain action and have reply
type Command interface {
	Actioner
	Replier
}

// State : An interface for certain state when user adding job that were able to reply, parse, do certain action, and proceed to next state
type State interface {
	Replier
	Parser
	Actioner
	Next
}
