package handler

import (
	iql "infoqerja-line/app/line"
	service "infoqerja-line/app/service"
	"log"
)

// Command : Interface for Reply service
type Command interface {
	Reply(bot iql.BotClient, token string) error
}

// HandleCommand : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func (handler *Handler) HandleCommand(command Command) error {
	// exec method
	return command.Reply(handler.Bot, handler.Token)
}

// Handler : A Handler for containing the necessary dependencies for all messages
type Handler struct {
	Bot   iql.BotClient
	Token string
}

// HandleIncomingMessage : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingMessage(bot iql.BotClient, token string, message string) {
	handler := &Handler{
		Bot:   bot,
		Token: token,
	}
	if IsValidCommand(message) {
		command := GetCommand(message)
		switch command {
		case "help":
			// inject incoming help service to the client handler
			if err := handler.HandleCommand(&service.IncomingHelp{}); err != nil {
				log.Print(err)
			}
		case "add":
			if err := handler.HandleCommand(&service.IncomingAdd{}); err != nil {
				log.Print(err)
			}
		case "show":
			if err := handler.HandleCommand(&service.IncomingShow{}); err != nil {
				log.Print(err)
			}
		default:
			// inject invalid service to the client handler
			if err := handler.HandleCommand(&service.IncomingInvalid{}); err != nil {
				log.Print(err)
			}
		}
	} else {
		// inject unknown service to the client handler
		if err := handler.HandleCommand(&service.IncomingUnknown{}); err != nil {
			log.Print(err)
		}
	}
}
