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

// IsValidCommand : Function to check wether user inputs is a command or not
func IsValidCommand(message string) bool {
	re := regexp.MustCompile("^!")
	return re.FindString(message) != ""
}

// GetCommand : get the type of command from user inputs
func GetCommand(command string) Command {
	co := strings.TrimSpace(command)
	if IsValidCommand(co) {
		switch command {
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
