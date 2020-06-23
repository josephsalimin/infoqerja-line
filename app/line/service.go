package line

import (
	iqq "infoqerja-line/app/command"
	"infoqerja-line/app/event"
	"infoqerja-line/app/model"
	state "infoqerja-line/app/state"
	"infoqerja-line/app/utils/constant"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type (
	// FinderCommand : interface of searching command service
	FinderCommand interface {
		GetCommand() model.Command
	}
	// CommandConstant : A service for searching something
	CommandConstant struct {
		Command string
	}
)

// GetCommand : get the type of command from user inputs
func (finder *CommandConstant) GetCommand() model.Command {
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
	// FinderState : interface of searching job service
	FinderState interface {
		GetState() model.State
	}
	// StateConstant : struct representing current state of user
	StateConstant struct {
		State string
	}
)

// GetState : get the type of state when waiting for user inputs
func (job *StateConstant) GetState() model.State {
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

type (
	// FinderEvent : interface for searching event service
	FinderEvent interface {
		GetEvent() model.Event
	}

	// EventConstant : struct representing current event requested by user
	EventConstant struct {
		Event string
	}
)

func (eve *EventConstant) GetEvent() model.Event {
	switch eve.Event {
	case constant.DetailEvent:
		return &event.Detail{}
	case constant.CancelEvent:
		return &event.Cancel{}
	case constant.StatsEvent:
		return &event.Stats{}
	default:
		return nil
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

// EventHandler : interface for injecting event service
type EventHandler interface {
	EventService(event model.Event) error
}

// CommandService : Method service for Command instance; the service that were going to be injected is the Command interface service
func (service *Service) CommandService(command model.Command) error {
	state, err := command.GetState()
	if state != nil {
		service.InputService(state)
	}
	_, err = service.Bot.ReplyMessage(service.Event.ReplyToken, command.GetReply()...).Do()
	return err
}

// InputService : Method service for State instance; the service that were going to be injected is the State interface service
func (service *Service) InputService(state model.State) error {
	if err := state.Parse(service.Event); err != nil {
		log.Print(err)
		return err
	}
	if err := state.Process(); err != nil {
		log.Print(err)
		return err
	}
	reply := state.GetReply()
	if err := state.NextState(); err != nil {
		log.Print(err)
		return err
	}
	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, reply...).Do()
	return err
}

// EventService : Method service for Event instance; the service that werge going to be injected is the Event interface service
func (service *Service) EventService(event model.Event) error {
	if err := event.Parse(service.Event); err != nil {
		log.Print(err)
		return err
	}

	if err := event.Process(); err != nil {
		log.Print(err)
		return err
	}

	_, err := service.Bot.ReplyMessage(service.Event.ReplyToken, event.GetReply()...).Do()
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
		finderLocal := &StateConstant{
			State: constant.Error,
		}
		errJob := finderLocal.GetState() // handling error
		_ = service.InputService(errJob)
	}
}

// HandleIncomingEvent : Handler for any incoming event that based on EventTypeMessage and EventTypePostback
func HandleIncomingEvent(service EventHandler, finder FinderEvent) {
	event := finder.GetEvent()
	if event != nil {
		if err := service.EventService(event); err != nil {
			log.Print(err)
		}
	}
}
