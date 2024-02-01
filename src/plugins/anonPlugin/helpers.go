package anonPlugin

import (
	"time"

	"github.com/ALiwoto/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

// LoadHandlers helper function will load all handlers for the current plugin.
func LoadHandlers(d *ext.Dispatcher, t []rune) {
	anonMessageHandler := handlers.NewMessage(anonMessageFilter, anonMessageHandler)
	blockCmd := handlers.NewCommand(blockCommand, blockCommandHandler)
	blockReqCallBack := handlers.NewCallback(blockReqCallBackQuery, blockReqBtnResponse)

	d.AddHandler(anonMessageHandler)
	d.AddHandler(blockCmd)
	d.AddHandler(blockReqCallBack)
}

func _getBlockRequestMap() *ssg.SafeEMap[int64, blockUserRequest] {
	m := ssg.NewSafeEMap[int64, blockUserRequest]()

	m.SetInterval(10 * time.Minute)
	m.SetExpiration(20 * time.Minute)

	return m
}
