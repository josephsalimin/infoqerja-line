package line

import (
	"infoqerja-line/app/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IncomingHelp : A class to represent the help command
type IncomingHelp struct{}

// Reply : Method service for IncomingHelp instance
func (handler *IncomingHelp) Reply(bot BotClient, token string) error {
	_, err := bot.ReplyMessage(token, linebot.NewTextMessage(constant.HelpMessage)).Do()
	return err
}
