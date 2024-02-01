package anonPlugin

import (
	"github.com/ALiwoto/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func (r *blockUserRequest) DeleteMessage() {
	if r.botMessage != nil {
		_, _ = r.botMessage.Delete(r.bot, nil)
	}
}

func (r *blockUserRequest) GetButtons() *gotgbot.InlineKeyboardMarkup {
	return &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text: "Confirm",
					CallbackData: blockReqCBData + sepChar +
						ssg.ToBase10(r.ownerId) + sepChar + blockReqConfirmCBData,
				},
				{
					Text: "✖️",
					CallbackData: blockReqCBData + sepChar +
						ssg.ToBase10(r.ownerId) + sepChar + blockReqDenyCBData,
				},
			},
		},
	}
}
