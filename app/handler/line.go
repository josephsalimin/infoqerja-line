package handler

import (
	iqc "infoqerja-line/app/config"
	iql "infoqerja-line/app/line"
	model "infoqerja-line/app/model"
	"infoqerja-line/app/utils"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"net/http"

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

					// addition : for adding job, we create new State for the user
					if message.Text == constant.AddCommandCode {
						customJobHandler(service, constant.NoState, utils.GetSource(*event), "")
					}
				} else {
					if user, err := (&util.UserDataReader{}).ReadOne(bson.M{
						constant.SourceID: utils.GetSource(*event),
					}); err == nil {
						customJobHandler(service, user.State, utils.GetSource(*event), message.Text)
					}
				}
			}
		case linebot.EventTypeFollow:
			customCommandHandler(service, constant.WelcomeCommandCode)
		case linebot.EventTypeUnfollow:
			customCommandHandler(service, constant.UnWelcomeCommandCode)
		case linebot.EventTypePostback:
			// checking user data -> get the state, and then verify it, create the CurrState struct data -> input into job sevice, check error, etc :)
			if user, err := (&util.UserDataReader{}).ReadOne(bson.M{
				constant.SourceID: utils.GetSource(*event),
			}); err == nil {
				postback := event.Postback.Data
				if postback == "DATE" {
					customJobHandler(service, user.State, utils.GetSource(*event), event.Postback.Params.Date)
				} else { // wrong input data
					customJobHandler(service, "error", utils.GetSource(*event), "")
				}
			}
		}
	}
}

// Private Method
func customCommandHandler(service *iql.Service, text string) {
	finder := &iql.Finder{
		Command: text,
	}
	iql.HandleIncomingCommand(service, finder)
}

// Private Method
func customJobHandler(service *iql.Service, currState, source, input string) {
	finder := &iql.JobState{
		State: currState,
	}
	data := model.BaseData{
		SourceID: source,
		Input:    input,
	}
	iql.HandleIncomingJob(service, finder, data)
}
