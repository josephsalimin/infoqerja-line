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
	// Command interface {
	// 	GetReply() []linebot.SendingMessage
	// }

	// FinderCommand : interface of searching command service
	FinderCommand interface {
		GetCommand() model.Command
	}

	// Finder : A service for searching something
	Finder struct {
		Command string
	}
)

// GetCommand : get the type of command from user inputs
func (finder *Finder) GetCommand() model.Command {
	co := strings.TrimSpace(finder.Command)
	switch co {
	case constant.HelpCommandCode:
		return &iqq.Help{}
	case constant.AddCommandCode:
		return &iqq.Add{}
	case constant.ShowCommandCode:
		return &iqq.Show{}
	case constant.WelcomeCommandCode: // hard coded code, for retrieving the welcome home page
		return &iqq.Welcome{}
	case constant.UnWelcomeCommandCode:
		return &iqq.UnWelcome{}
	default:
		return &iqq.Invalid{}
	}
}

type (
	// Job : Interface for Executing a job
	// Job interface {
	// 	Execute() error
	// 	// make 4 method
	// }

	// FinderJob : interface of searching job service
	FinderState interface {
		GetState() model.State
	}

	// JobState : struct representing current state of user
	JobState struct {
		State string
	}
)

// GetJob : get the type of command from user inputs
func (job *JobState) GetState() model.State {
	switch job.State {
	case constant.WaitDateInput:
		return &state.AddDateState{}
	case constant.WaitDescInput:
		return &state.AddDescState{}
	case constant.WaitTitleInput:
		return &state.AddTitleState{}
	case constant.NoState:
		return &state.StartState{}
	default:
		return &state.ErrorState{}
	}
}

// Service : A Handler for containing the necessary dependencies for all messages
type Service struct {
	Bot   BotClient
	Event linebot.Event
}

// Commander : interface for injecting messaging service
type Commander interface {
	CommandService(command model.Command) error
}

// Inputer : interface for injecting job service
type Inputer interface {
	InputService(state model.State) error
}

// CommandService : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
func (service *Service) CommandService(command model.Command) error {

	// execute the action
	state, err := command.GetState()
	if state != nil {
		service.InputService(state)
	}

	// reply
	_, err = service.Bot.ReplyMessage(service.Event.ReplyToken, command.GetReply()...).Do()
	return err
}

// InputService : Method service for IncomingJob instance; the service that were going to be injected is the Job interface service
func (service *Service) InputService(state model.State) error {
	// parse
	if err := state.Parse(service.Event); err != nil {
		log.Print(err)
		return err
	}
	// process
	if err := state.Process(); err != nil {
		log.Print(err)
		return err
	}
	// get reply
	reply := state.GetReply()
	// next state
	if err := state.NextState(); err != nil {
		log.Print(err)
		return err
	}
	// bales reply :)
	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, reply...).Do()
	return err
}

// HandleIncomingCommand : Handler for any incoming event that based on EventTypeMessage
func HandleIncomingCommand(service Commander, finder FinderCommand) {
	command := finder.GetCommand()
	if command != nil {
		if err := service.CommandService(command); err != nil {
			log.Print(err)
		}
	}
}

// HandleIncomingJob : Handler for any incoming job that based on EventTypeMessage and EventTypePostback
func HandleIncomingJob(service Inputer, finder FinderState) {
	job := finder.GetState()
	if err := service.InputService(job); err != nil {
		log.Print(err)
		finderLocal := &JobState{
			State: "error",
		}
		errJob := finderLocal.GetState() // handling error
		_ = service.InputService(errJob)
	}
}
