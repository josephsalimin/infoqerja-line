package main

import (
	"infoqerja-line/app/config"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
)

func main() {
	configService, err := config.ReadConfig(nil)

	if err != nil {
		log.Fatal(err)
	}

	config := configService.GetConfig()
	addr := config.Host + ":" + strconv.Itoa(config.Port)

	r := GetApplicationRouter(*config)
	router := handlers.LoggingHandler(os.Stdout, r)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
