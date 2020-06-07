package handler

import (
	iql "infoqerja-line/app/line"
	"log"
)

// Command : Interface for Reply service
type Command interface {
	Reply(bot iql.BotClient, token string) error
}

// HandleCommand : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func HandleCommand(command Command, bot iql.BotClient, token string) error {
	// exec method
	return command.Reply(bot, token)
}

// HandleIncomingMessage : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingMessage(bot iql.BotClient, token string, message string) {
	if IsValidCommand(message) {
		command := GetCommand(message)
		switch command {
		case "help":
			// inject incoming help service to the client handler
			if err := HandleCommand(&iql.IncomingHelp{}, bot, token); err != nil {
				log.Print(err)
			}
		default:
			// inject invalid service to the client handler
			if err := HandleCommand(&iql.IncomingInvalid{}, bot, token); err != nil {
				log.Print(err)
			}
		}
	} else {
		// inject unknown service to the client handler
		if err := HandleCommand(&iql.IncomingUnknown{}, bot, token); err != nil {
			log.Print(err)
		}
	}
}
