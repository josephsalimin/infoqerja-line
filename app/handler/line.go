package handler

import (
	iqc "infoqerja-line/app/config"
	crud "infoqerja-line/app/crud"
	iql "infoqerja-line/app/line"
	"infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
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
				if iql.IsValidCommand(message.Text) {
					finder := &iql.Finder{
						Command: message.Text,
					}
					iql.HandleIncomingCommand(service, finder)

					// suggesstion : create more sophisticated if about this edge of code
					if message.Text == constant.AddCommandCode {
						finder := &iql.JobState{
							State: constant.NoState, // for adding new user to the service batch database
						}
						iql.HandleIncomingService(service, finder)
					}
				} else {
					// read user data
					user, err := crud.ReadSingleUserData(utils.GetSource(*event))
					// only if the user is recognized as an applicant of inserting data
					if err == nil {
						finder := &iql.JobState{
							State: user.State,
						}
						// handle the incoming service
						iql.HandleIncomingService(service, finder)
					} // means it just a normal chat
				}
			}
		case linebot.EventTypeFollow:
			// add welcome handler
			finder := &iql.Finder{
				Command: constant.WelcomeCommandCode,
			}
			iql.HandleIncomingCommand(service, finder)
		case linebot.EventTypeUnfollow:
			finder := &iql.Finder{
				Command: constant.UnWelcomeCommandCode,
			}
			iql.HandleIncomingCommand(service, finder)
		case linebot.EventTypePostback:
			// checking user data -> get the state, and then verify it, create the CurrState struct data -> input into job sevice, check error, etc :)
			user, err := crud.ReadSingleUserData(utils.GetSource(*event))
			// only if the user is recognized as an applicant of inserting data
			if err == nil {
				finder := &iql.JobState{
					State: user.State,
				}
				// handle the incoming service
				iql.HandleIncomingService(service, finder)
			} // means it just a normal chat

			// getting the data --> for the date data
			// data := event.Postback.Data
			// if data == "DATE" {
			// 	log.Printf("Successful getting data : (%v)", *&event.Postback.Params.Date)
			// }
			// finder := &util.Finder{
			// 	Command: "!show",
			// }
			// iql.HandleIncomingCommand(service, finder)
		}
	}
}
