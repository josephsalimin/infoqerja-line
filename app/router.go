package main

import (
	iqc "infoqerja-line/app/config"
	handler "infoqerja-line/app/handler"
	iql "infoqerja-line/app/line"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(config iqc.Config, bot iql.BotClient) *mux.Router {
	r := mux.NewRouter()
	indexHandler := handler.BuildIndexHandler(config)
	lineBotHandler := handler.BuildLineBotHandler(config, bot)

	r.HandleFunc("/", indexHandler.Welcome)
	r.HandleFunc("/line", lineBotHandler.Callback)

	return r
}
