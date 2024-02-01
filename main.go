package main

import (
	"log"

	"github.com/ALiwoto/AnonimasuRobot/src/core/logging"
	"github.com/ALiwoto/AnonimasuRobot/src/core/wotoConfig"
	"github.com/ALiwoto/AnonimasuRobot/src/database"
	"github.com/ALiwoto/AnonimasuRobot/src/plugins"
)

func main() {
	_, err := wotoConfig.LoadConfig()
	if err != nil {
		log.Fatal("Error parsing config file", err)
	}

	f := logging.LoadLogger()
	if f != nil {
		defer f()
	}

	err = database.StartDB()
	if err != nil {
		logging.Fatal("Failed to start database: ", err)
	}

	err = plugins.StartTelegramBot()
	if err != nil {
		logging.Fatal("Failed to start the bot: ", err)
	}
}
