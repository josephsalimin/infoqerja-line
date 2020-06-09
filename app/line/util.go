package line

import (
	"regexp"
	"strings"

	iqq "infoqerja-line/app/command"
	constant "infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Command : Interface for Reply service
type Command interface {
	GetMessage() []linebot.SendingMessage
}

// FinderCommand : interface of searching command service
type FinderCommand interface {
	GetCommand() Command
}

// Finder : A service for searching something
type Finder struct {
	Command string
}

// IsValidCommand : Function to check wether user inputs is a command or not
func IsValidCommand(message string) bool {
	re := regexp.MustCompile("^!")
	return re.FindString(message) != ""
}

// GetCommand : get the type of command from user inputs
func (finder *Finder) GetCommand() Command {
	co := strings.TrimSpace(finder.Command)
	if IsValidCommand(co) {
		switch finder.Command {
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
	} else {
		return nil
	}
}
