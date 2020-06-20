package utils

import (
	"regexp"
	"strings"

	iqq "infoqerja-line/app/event/command"
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

type (
	// Command : Interface for Reply service
	Command interface {
		GetReply() []linebot.SendingMessage
	}

	// FinderCommand : interface of searching command service
	FinderCommand interface {
		GetCommand() Command
	}

	// Finder : A service for searching something
	Finder struct {
		Command string
	}
)

// IsValidCommand : Function to check wether user inputs is a command or not
func IsValidCommand(message string) bool {
	re := regexp.MustCompile("^!")
	return re.FindString(message) != ""
}

// GetCommand : get the type of command from user inputs
func (finder *Finder) GetCommand() Command {
	co := strings.TrimSpace(finder.Command)
	switch co {
	case constant.HelpCommandCode:
		return &iqq.IncomingHelp{}
	case constant.AddCommandCode:
		return &iqq.IncomingAdd{}
	case constant.ShowCommandCode:
		return &iqq.IncomingShow{}
	case constant.WelcomeCommandCode: // hard coded code, for retrieving the welcome home page
		return &iqq.Welcome{}
	case constant.UnWelcomeCommandCode:
		return &iqq.UnWelcome{}
	default:
		return &iqq.IncomingInvalid{}
	}
}
