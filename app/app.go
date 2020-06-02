package main

import (
	"infoqerja-line/app/config"
	"infoqerja-line/app/line"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
)

func initializeConfig() *config.Service {
	configService, err := config.ReadConfig(&config.EnvReader{})

	if err != nil {
		log.Fatal(err)
	}

	return configService
}

func initializeLineBot(config config.Config) line.BotClient {
	bot, err := line.InitializeBot(config)

	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func main() {
	configService := initializeConfig()
	config := *configService.GetConfig()
	bot := initializeLineBot(config)

	addr := config.Host + ":" + strconv.Itoa(config.Port)

	r := GetApplicationRouter(config, bot)
	router := handlers.LoggingHandler(os.Stdout, r)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
