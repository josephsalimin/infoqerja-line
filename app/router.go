package main

import (
	iqc "infoqerja-line/app/config"
	handler "infoqerja-line/app/handler"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(config iqc.Config) *mux.Router {
	r := mux.NewRouter()
	indexHandler := handler.IndexHandler{Config: config}

	r.HandleFunc("/", indexHandler.Welcome)

	return r
}
