package handler

import (
	bot "infoqerja-line/app/bot"
	iqc "infoqerja-line/app/config"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BotHandler will handle all bot's callback request
type BotHandler struct {
	Config iqc.Config
	Bot    bot.LineBotService
}

// HandleMessage will handle the callback from line
func (h BotHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	events, err := h.Bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = h.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
