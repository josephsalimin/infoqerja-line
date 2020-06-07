package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Service : A Handler for containing the necessary dependencies for all messages
type Service struct {
	Bot   BotClient
	Event linebot.Event
}

// MessageService : interface for injecting messaging service
type MessageService interface {
	MessageServiceReply(command Command) error
}

// MessageServiceReply : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func (service *Service) MessageServiceReply(command Command) error {
	// exec methoda
	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, command.GetMessage()...).Do()
	return err
}

// HandleIncomingMessage : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingMessage(service MessageService, message string) {
	// get command
	command := GetCommand(message)
	// exec something
	// reply
	if command != nil {
		if err := service.MessageServiceReply(command); err != nil {
			log.Print(err)
		}
	}
}
