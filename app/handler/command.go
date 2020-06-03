package handler

import (
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

const welcomeMessage = `Welcome to the InfoQerja Bot!!!💻💻`
const helpMessage = `Use command below to use InfoQerja functionality:
- !help		: to find out how to use InfoQerja
- !add		: to add job posting to InfoQerja
- !show		: to show job posting in InfoQerja
`

const invalidMessage = `Please enter a valid command!! Refer to !help for available command.`

const unknownMessage = `This bot does not respond to other things except for command!! 😎😎
Please refer to !help command to use InfoQerja functionality.
Hope you enjoy this bot !!😊😊
- Joseph Salimin 😍
`

// IsValidCommand : Function to check wether user inputs is a command or not
func IsValidCommand(message string) bool {
	re := regexp.MustCompile("^!")
	return re.FindString(message) != ""
}

// CheckCommand : get the type of command from user inputs
func CheckCommand(command string) string {
	co := strings.TrimSpace(command)
	return co[1:]
}

// HelpCommand : Function to respond for the !help command
func (h LineBotHandler) HelpCommand(token string) error {
	_, err := h.bot.ReplyMessage(token, linebot.NewTextMessage(helpMessage)).Do()
	return err
}

// AddCommand : Function to respond for the !add command
func (h LineBotHandler) AddCommand(token string) {
	// not yet implemented
}

// ShowCommand : Function to respond for the !show command
func ShowCommand() {
	// not yet implemented
}

// InvalidCommand : Function to respond for the !<invalid> command. To give better instruction
func (h LineBotHandler) InvalidCommand(token string) error {
	_, err := h.bot.ReplyMessage(token, linebot.NewTextMessage(invalidMessage)).Do()
	return err
}

// WelcomeHandler : Respond to follow event
func (h LineBotHandler) WelcomeHandler(token string) error {
	_, err := h.bot.ReplyMessage(token, linebot.NewTextMessage(welcomeMessage+"\n"+helpMessage)).Do()
	return err
}
