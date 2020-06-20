package line

import (
	util "infoqerja-line/app/utils"
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
	MessageServiceReply(command util.Command) error
}

// JobService : interface for injecting job service
type JobService interface {
	ExecuteJob() error
}

// MessageServiceReply : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func (service *Service) MessageServiceReply(command util.Command) error {
	// exec methoda
	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, command.GetReply()...).Do()
	return err
}

// JobServiceExecute : Method service for IncomingJob instance; the service that were going to be injected is the Job interface service
func (service *Service) JobServiceExecute(job util.Job) error {
	// executing the method
	if err := job.Execute(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// HandleIncomingCommand : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingCommand(service MessageService, finder util.FinderCommand) {
	command := finder.GetCommand()
	if command != nil {
		if err := service.MessageServiceReply(command); err != nil {
			log.Print(err)
		}
	}
}

func HandleIncomingService(service JobService, finder util.FinderCommand) {

}

// GetSource : Get source for any event happening to bot
func GetSource(event linebot.Event) string {
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		return event.Source.UserID
	case linebot.EventSourceTypeGroup:
		return event.Source.UserID
	case linebot.EventSourceTypeRoom:
		return event.Source.RoomID
	}
	return event.Source.UserID
}

