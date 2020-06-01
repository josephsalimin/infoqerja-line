package main

import (
	"infoqerja-line/app/bot"
	"infoqerja-line/app/config"
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

func initializeBot(config config.Config) bot.LineBotService {
	bot, err := bot.InitializeLineBot(config)

	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func main() {
	configService := initializeConfig()
	config := *configService.GetConfig()
	bot := initializeBot(config)

	addr := config.Host + ":" + strconv.Itoa(config.Port)

	r := GetApplicationRouter(config, bot)
	router := handlers.LoggingHandler(os.Stdout, r)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
