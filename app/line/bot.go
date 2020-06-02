package line

import (
	"infoqerja-line/app/config"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BotClient mimics line-bot-sdk-go Client
type BotClient interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) BotPushMessageCall
}

// BotPushMessageCall mimics line-bot-sdk-go ReplyMessageCall
type BotPushMessageCall interface {
	Do() (*linebot.BasicResponse, error)
}

// InfoQerjaBot is bot implementation that contains actual line-bot-sdk-go
type InfoQerjaBot struct {
	bot *linebot.Client
}

// ParseRequest will call line-bot-sdk-go client's ParseRequest
func (iqb InfoQerjaBot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return iqb.bot.ParseRequest(r)
}

// ReplyMessage will call line-bot-sdk-go client's ReplyMessage
func (iqb InfoQerjaBot) ReplyMessage(replyToken string, messages ...linebot.SendingMessage) BotPushMessageCall {
	return iqb.bot.ReplyMessage(replyToken, messages...)
}

// InitializeBot initiate line-bot-sdk-go client
func InitializeBot(config config.Config) (BotClient, error) {
	bot, err := linebot.New(
		config.ChannelSecret,
		config.ChannelToken,
	)

	if err != nil {
		return nil, err
	}

	return InfoQerjaBot{bot: bot}, nil
}
