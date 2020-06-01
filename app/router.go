package main

import (
	iqb "infoqerja-line/app/bot"
	iqc "infoqerja-line/app/config"
	handler "infoqerja-line/app/handler"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(config iqc.Config, bot iqb.LineBotService) *mux.Router {
	r := mux.NewRouter()
	indexHandler := handler.IndexHandler{Config: config}
	botHandler := handler.BotHandler{Config: config, Bot: bot}

	r.HandleFunc("/", indexHandler.Welcome)
	r.HandleFunc("/bot", botHandler.HandleMessage)

	return r
}
