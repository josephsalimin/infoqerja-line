package handler

import (
	iqc "infoqerja-line/app/config"
	constant "infoqerja-line/app/constant"
	iql "infoqerja-line/app/line"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotHandler will handle all line's callback request
type LineBotHandler struct {
	config iqc.Config
	bot    iql.BotClient
}

// BuildLineBotHandler returns LineBotHandler struct
func BuildLineBotHandler(config iqc.Config, bot iql.BotClient) *LineBotHandler {
	return &LineBotHandler{
		config: config,
		bot:    bot,
	}
}

// Callback will handle the callback from line
func (h LineBotHandler) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := h.bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		service := &iql.Service{
			Bot:   h.bot,
			Event: *event,
		}

		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				iql.HandleIncomingMessage(service, message.Text)
			}
		case linebot.EventTypeFollow:
			// add welcome handler
			iql.HandleIncomingMessage(service, constant.WelcomeCommandCode)
		case linebot.EventTypeUnfollow:
			// add welcome handler
			iql.HandleIncomingMessage(service, constant.UnWelcomeCommandCode)
		case linebot.EventTypeJoin:
			iql.HandleIncomingMessage(service, constant.WelcomeCommandCode)
		case linebot.EventTypeLeave:
			iql.HandleIncomingMessage(service, constant.UnWelcomeCommandCode)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" {
				log.Printf("Successful getting data : (%v)", *&event.Postback.Params.Date)
			}
			iql.HandleIncomingMessage(service, "!show")
		}
	}
}
