package plugins

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ALiwoto/AnonimasuRobot/src/core/logging"
	"github.com/ALiwoto/AnonimasuRobot/src/core/wotoConfig"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartTelegramBot() error {
	token := wotoConfig.GetBotToken()
	if len(token) == 0 {
		return errors.New("bot token is empty")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: 6 * gotgbot.DefaultTimeout,
		},
	})
	if err != nil {
		return err
	}

	updater := ext.NewUpdater(nil)
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
	})
	if err != nil {
		return err
	}

	logging.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))

	LoadAllHandlers(updater.Dispatcher, wotoConfig.GetCmdPrefixes())

	updater.Idle()
	return nil
}
