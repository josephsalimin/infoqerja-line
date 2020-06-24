package handler

import (
	iqc "infoqerja-line/app/config"
	iql "infoqerja-line/app/line"
	"infoqerja-line/app/utils"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
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
				if util.IsCommandValid(message.Text) {
					customCommandHandler(service, message.Text)
					// add else if -> is valid event; do the event using customEventHandler, for future upgrade on the code
				} else {
					if user, err := (&util.UserDataReader{}).ReadOne(bson.M{
						constant.SourceID: utils.GetSource(*event),
						constant.State: bson.M{
							"$in": bson.A{constant.WaitTitleInput, constant.WaitDescInput},
						},
					}); err == nil && user != nil {
						customJobHandler(service, user.State)
					}
				}
			}
		case linebot.EventTypeFollow:
			customCommandHandler(service, constant.WelcomeCommandCode)
		case linebot.EventTypeUnfollow:
			customCommandHandler(service, constant.UnWelcomeCommandCode)
		case linebot.EventTypePostback:
			// checking user data -> get the state, and then verify it, create the CurrState struct data -> input into job sevice, check error, etc :)
			postback := event.Postback.Data
			if postback == constant.DateData {
				if user, err := (&util.UserDataReader{}).ReadOne(bson.M{
					constant.SourceID: utils.GetSource(*event),
					constant.State:    constant.WaitDateInput,
				}); err == nil && user != nil {
					customJobHandler(service, user.State)
				} else {
					customJobHandler(service, constant.Error)
				}
			} else if strings.Contains(postback, constant.JobIDData) {
				customEventHandler(service, constant.DetailEvent)
			}
		}
	}
}

// Private Method
func customCommandHandler(service *iql.Service, text string) {
	finder := &iql.CommandConstant{
		Command: text,
	}
	iql.HandleIncomingCommand(service, finder)
}

// Private Method
func customJobHandler(service *iql.Service, currState string) {
	finder := &iql.StateConstant{
		State: currState,
	}
	iql.HandleIncomingJob(service, finder)
}

func customEventHandler(service *iql.Service, event string) {
	finder := &iql.EventConstant{
		Event: event,
	}
	iql.HandleIncomingEvent(service, finder)
}
