package line

// import (
// 	"log"
// 	"regexp"
// 	"strings"
// )

// // Command : Interface for Reply service
// type Command interface {
// 	Reply(bot BotClient, token string) error
// }

// // HandleCommand : Method service for IncomingAction instance; the service that were going to be injected is the Command interface service
// func HandleCommand(command Command, bot BotClient, token string) error {
// 	// exec method
// 	return command.Reply(bot, token)
// }

// // IsValidCommand : Function to check wether user inputs is a command or not
// func IsValidCommand(message string) bool {
// 	re := regexp.MustCompile("^!")
// 	return re.FindString(message) != ""
// }

// // GetCommand : get the type of command from user inputs
// func GetCommand(command string) string {
// 	co := strings.TrimSpace(command)
// 	return co[1:]
// }

// // HandleIncomingMessage : Handler for any incoming event that based on EventTypeMessage
// func HandleIncomingMessage(bot BotClient, token string, message string) {
// 	if IsValidCommand(message) {
// 		command := GetCommand(message)
// 		switch command {
// 		case "help":
// 			// inject incoming help service to the client handler
// 			if err := HandleCommand(&IncomingHelp{}, bot, token); err != nil {
// 				log.Print(err)
// 			}
// 		default:
// 			// inject invalid service to the client handler
// 			if err := HandleCommand(&IncomingInvalid{}, bot, token); err != nil {
// 				log.Print(err)
// 			}
// 		}
// 	} else {
// 		// inject unknown service to the client handler
// 		if err := HandleCommand(&IncomingUnknown{}, bot, token); err != nil {
// 			log.Print(err)
// 		}
// 	}
// }
