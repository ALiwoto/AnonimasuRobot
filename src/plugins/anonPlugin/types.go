package anonPlugin

import "github.com/PaulSonOfLars/gotgbot/v2"

type blockUserRequest struct {
	// ownerId is the owner of the request.
	ownerId int64

	// targetId is the target user's id to be blocked.
	targetId int64

	botMessage *gotgbot.Message
	bot        *gotgbot.Bot
}
