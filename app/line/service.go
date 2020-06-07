package line

import (
	"log"
)

// MessageService : A Handler for containing the necessary dependencies for all messages
type MessageService struct {
	Bot   BotClient
	Token string
}

// MessageServiceReply : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func MessageServiceReply(service MessageService, command Command) error {
	// exec methoda
	_, err := service.Bot.ReplyMessage(service.Token, command.GetMessage()...).Do()
	return err
}

// HandleIncomingMessage : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingMessage(handler MessageService, message string) {
	command := GetCommand(message)
	if err := MessageServiceReply(handler, command); err != nil {
		log.Print(err)
	}
}
