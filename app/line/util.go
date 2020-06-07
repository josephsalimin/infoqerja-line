package line

import (
	"regexp"
	"strings"

	iqq "infoqerja-line/app/command"

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
		case "!help":
			return &iqq.IncomingHelp{}
		case "!add":
			return &iqq.IncomingAdd{}
		case "!show":
			return &iqq.IncomingShow{}
		default:
			return &iqq.IncomingInvalid{}
		}
	} else {
		return &iqq.IncomingUnknown{}
	}
}
