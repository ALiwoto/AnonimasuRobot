package database

import (
	"sync"

	"github.com/AnimeKaizoku/AnonimasuRobot/src/core/logging"
	wv "github.com/AnimeKaizoku/AnonimasuRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/AnonimasuRobot/src/database/usersDatabase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartDB() error {
	var err error
	var db *gorm.DB

	conf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err = gorm.Open(sqlite.Open("./Database/AnonimasuRobot.db"), conf)
	if err != nil {
		return err
	}

	logging.Info("Database connected ")

	//Create tables if they don't exist
	err = wv.Core.AutoMigrateDB(
		usersDatabase.ModelBlockedUser,
	)
	if err != nil {
		return err
	}

	logging.Info("Auto-migrated database schema")

	wv.Core.DB = db
	wv.Core.DbMutex = &sync.Mutex{}

	return nil
}
