package line

import "github.com/line/line-bot-sdk-go/linebot"

const welcomeMessage = `Welcome to the _InfoQerja Bot!!!_💻💻`
const helpMessage = `Use command below to use InfoQerja functionality:
- *!help*		: to find out how to use InfoQerja
- *!add*		: to add job posting to InfoQerja
- *!show*		: to show job posting in InfoQerja
`

const invalidMessage = `Please enter a valid command!! Refer to *!help* for available command.`

const unknownMessage = `This bot does not respond to other things except for command!! 😎😎
Please refer to *!help* command to use InfoQerja functionality.
Hope you enjoy this bot !!😊😊
- Joseph Salimin 😍
`

// Action : Interface for Reply service
type Action interface {
	Reply(token string) error
}

// IncomingAction : Instance for Handling Incoming Message
type IncomingAction struct {
	Replier Action
}

// IncomingHelp :  Instance for handling help message
type IncomingHelp struct {
	Bot BotClient
}

// IncomingUnknown :  Instance for handling unknown message
type IncomingUnknown struct {
	Bot BotClient
}

// IncomingInvalid :  Instance for handling invalid message
type IncomingInvalid struct {
	Bot BotClient
}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingHelp) Reply(token string) error {
	_, err := handler.Bot.ReplyMessage(token, linebot.NewTextMessage(helpMessage)).Do()
	return err
}

// Reply : Method service for IncomingInvalid instance
func (handler *IncomingInvalid) Reply(token string) error {
	_, err := handler.Bot.ReplyMessage(token, linebot.NewTextMessage(invalidMessage)).Do()
	return err
}

// Reply : Method service for IncomingUnknown instance
func (handler *IncomingUnknown) Reply(token string) error {
	_, err := handler.Bot.ReplyMessage(token, linebot.NewStickerMessage("11539", "52114146"), linebot.NewTextMessage(unknownMessage)).Do()
	return err
}

// HandleIncomingMessage : Method service for IncomingAction instance
func (handler *IncomingAction) HandleIncomingMessage(token string) error {
	return handler.Replier.Reply(token)
}
