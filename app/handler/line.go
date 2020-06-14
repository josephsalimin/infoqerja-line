package handler

import (
	iqc "infoqerja-line/app/config"
	iql "infoqerja-line/app/line"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
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
				if util.IsValidCommand(message.Text) {
					finder := &util.Finder{
						Command: message.Text,
					}
					iql.HandleIncomingMessage(service, finder)
				} else {
					// handle incoming normal message
				}
				// default : don;t need to take care of this
			}
		case linebot.EventTypeFollow:
			// add welcome handler
			finder := &util.Finder{
				Command: constant.WelcomeCommandCode,
			}
			iql.HandleIncomingMessage(service, finder)
		case linebot.EventTypeUnfollow:
			finder := &util.Finder{
				Command: constant.UnWelcomeCommandCode,
			}
			iql.HandleIncomingMessage(service, finder)
		case linebot.EventTypePostback:

			// checking user data

			data := event.Postback.Data
			if data == "DATE" {
				log.Printf("Successful getting data : (%v)", *&event.Postback.Params.Date)
			}
			finder := &util.Finder{
				Command: "!show",
			}
			iql.HandleIncomingMessage(service, finder)

			
		}
	}
}
