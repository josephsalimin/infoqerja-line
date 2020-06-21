package model

import "github.com/line/line-bot-sdk-go/linebot"

// BaseData : the basec structure for data (DTO) in the event layer
type BaseData struct {
	SourceID string
	Input    string
}

type (
	Actioner interface {
		Do() error
	}

	Replier interface {
		GetReply() []linebot.SendingMessage
	}
	Parser interface {
		Parse() error
	}
	Next interface {
		NextState() error
	}
)

type (
	Command interface {
		Actioner
		Replier
	}
)

type (
	State interface {
		Replier
		Parser
		Actioner
		Next
	}
)
