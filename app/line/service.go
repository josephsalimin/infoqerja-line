package line

import (
	iqq "infoqerja-line/app/command"
	"infoqerja-line/app/model"
	state "infoqerja-line/app/state"
	"infoqerja-line/app/utils/constant"
	"log"
	"strings"

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

// GetCommand : get the type of command from user inputs
func (finder *Finder) GetCommand() Command {
	co := strings.TrimSpace(finder.Command)
	switch co {
	case constant.HelpCommandCode:
		return &iqq.IncomingHelp{}
	case constant.AddCommandCode:
		return &iqq.Add{}
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

type (
	// Job : Interface for Executing a job
	Job interface {
		Execute() error
		// make 4 method
	}

	// FinderJob : interface of searching job service
	FinderJob interface {
		GetJob(data model.BaseData) Job
	}

	// JobState : struct representing current state of user
	JobState struct {
		State string
	}
)

// GetJob : get the type of command from user inputs
func (job *JobState) GetJob(data model.BaseData) Job {
	switch job.State {
	case constant.WaitDateInput:
		return &state.IncomingAddDateJob{
			Data: data,
		}
	case constant.WaitDescInput:
		return &state.IncomingAddDescJob{
			Data: data,
		}
	case constant.WaitTitleInput:
		return &state.IncomingAddTitleJob{
			Data: data,
		}
	case constant.NoState:
		return &state.IncomingStartInput{
			Data: data,
		}
	default:
		return &state.IncomingErrorEvent{
			Data: data,
		}
	}
}

// Service : A Handler for containing the necessary dependencies for all messages
type Service struct {
	Bot   BotClient
	Event linebot.Event
}

// MessageService : interface for injecting messaging service
type MessageService interface {
	MessageServiceReply(command Command) error
}

// JobService : interface for injecting job service
type JobService interface {
	JobServiceExecute(job Job) error
}

// MessageServiceReply : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func (service *Service) MessageServiceReply(command Command) error {
	// exec methoda
	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, command.GetReply()...).Do()
	return err
}

// JobServiceExecute : Method service for IncomingJob instance; the service that were going to be injected is the Job interface service
func (service *Service) JobServiceExecute(job Job) error {
	// executing the method
	if err := job.Execute(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// HandleIncomingCommand : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingCommand(service MessageService, finder FinderCommand) {
	command := finder.GetCommand()
	if command != nil {
		if err := service.MessageServiceReply(command); err != nil {
			log.Print(err)
		}
	}
}

// HandleIncomingJob : Handler for any incoming job that based on EventTypeMessage and EventTypePostback
func HandleIncomingJob(service JobService, finder FinderJob, data model.BaseData) {
	job := finder.GetJob(data)
	// filling job description data
	if err := service.JobServiceExecute(job); err != nil {
		log.Print(err)
		finderLocal := &JobState{
			State: "error",
		}
		dataLocal := &model.BaseData{
			SourceID: data.SourceID,
		}
		errJob := finderLocal.GetJob(*dataLocal) // handling error
		_ = service.JobServiceExecute(errJob)
	}
}
