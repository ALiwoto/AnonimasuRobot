package plugins

import (
	"github.com/ALiwoto/AnonimasuRobot/src/plugins/anonPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	anonPlugin.LoadHandlers(d, triggers)
}
