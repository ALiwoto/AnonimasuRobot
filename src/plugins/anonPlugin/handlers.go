package anonPlugin

import (
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/ALiwoto/AnonimasuRobot/src/core/utils"
	"github.com/ALiwoto/AnonimasuRobot/src/core/wotoConfig"
	"github.com/ALiwoto/AnonimasuRobot/src/database/usersDatabase"
	"github.com/ALiwoto/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func anonMessageFilter(msg *gotgbot.Message) bool {
	return msg.Chat.Type == "private"
}

func anonMessageHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if usersDatabase.IsUserBlocked(user.Id) {
		return ext.EndGroups
	} else if ctx.Message.Text == "" || strings.HasPrefix(msg.Text, "/") {
		return ext.ContinueGroups
	}

	chatMember, err := bot.GetChatMember(wotoConfig.GetTargetChat(), user.Id, nil)
	if err != nil {
		return ext.EndGroups
	}

	switch chatMember.GetStatus() {
	case "left", "kicked":
		return ext.EndGroups
	}

	var messageId int64
	repliedId := utils.ExtractRepliedToMessageId(ctx.Message)
	if repliedId == 0 {
		message, err := ctx.Message.Copy(bot, wotoConfig.GetTargetChat(), &gotgbot.CopyMessageOpts{
			ReplyToMessageId: utils.ExtractRepliedToMessageId(ctx.Message),
		})
		if err != nil || message == nil {
			return ext.EndGroups
		}

		messageId = message.MessageId
	} else {
		message, err := bot.SendMessage(wotoConfig.GetTargetChat(), ctx.Message.Text, &gotgbot.SendMessageOpts{
			ReplyToMessageId: repliedId,
			Entities:         ctx.Message.Entities,
		})
		if err != nil || message == nil {
			return ext.EndGroups
		}

		messageId = message.MessageId
	}

	usersDatabase.SetUserFromMessageId(messageId, user)

	return ext.EndGroups
}

func blockCommandHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Id != wotoConfig.GetTargetChat() ||
		ctx.EffectiveMessage.ReplyToMessage == nil {
		return ext.ContinueGroups
	}

	user := ctx.EffectiveUser
	replied := ctx.EffectiveMessage.ReplyToMessage
	chatMember, err := bot.GetChatMember(wotoConfig.GetTargetChat(), user.Id, nil)
	if err != nil {
		return ext.EndGroups
	}

	switch chatMember.GetStatus() {
	case "owner", "administrator":
		break
	default:
		return ext.EndGroups
	}

	targetUser := usersDatabase.GetUserFromMessageId(replied.MessageId)
	if targetUser == 0 {
		md := mdparser.GetNormal("Looks like your /block command was sent to a very old message.")
		_, _ = bot.SendMessage(user.Id, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	previousReq := blockRequestMap.Get(user.Id)
	if previousReq != nil {
		previousReq.DeleteMessage()
	}

	blockReq := &blockUserRequest{
		ownerId:  user.Id,
		targetId: targetUser,
	}
	blockRequestMap.Add(blockReq.ownerId, blockReq)

	md := mdparser.GetNormal("you are about to block an anon user and this action will block them, whoever they are")
	md.Bold(" " + wotoConfig.GetBlockExpirationString() + ".")
	md.Normal("\nAre you sure you want to continue?")

	blockReq.botMessage, _ = ctx.EffectiveMessage.Reply(bot, md.ToString(), &gotgbot.SendMessageOpts{
		ReplyMarkup: blockReq.GetButtons(),
		ParseMode:   gotgbot.ParseModeMarkdownV2,
	})
	blockReq.bot = bot

	return ext.EndGroups
}

func blockReqCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, blockReqCBData+sepChar)
}

func blockReqBtnResponse(bot *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.CallbackQuery
	allStrs := ssg.Split(query.Data, sepChar)
	msg := query.Message
	// format is blockReq_ownerId_[y/n]
	if len(allStrs) < 3 {
		return ext.EndGroups
	}

	ownerId := ssg.ToInt64(allStrs[1])
	if ownerId != query.From.Id {
		_, _ = query.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "This button is not for you!",
			ShowAlert: true,
			CacheTime: 5800,
		})
		return ext.EndGroups
	}

	blockReq := blockRequestMap.Get(ownerId)
	if blockReq == nil {
		_, _ = query.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "You are too late! Send block command again.",
			ShowAlert: true,
			CacheTime: 5800,
		})
		_, _ = msg.Delete(bot, nil)
		return ext.EndGroups
	}

	switch allStrs[2] {
	case blockReqDenyCBData:
		blockReq.DeleteMessage()
		return ext.EndGroups
	case blockReqConfirmCBData:
		usersDatabase.BlockUser(blockReq.targetId)
		_, _, _ = msg.EditText(bot, mdparser.GetMono("Anon user has been blocked.").ToString(),
			&gotgbot.EditMessageTextOpts{
				ParseMode: gotgbot.ParseModeMarkdownV2,
			})
	default:
		return ext.EndGroups
	}

	return ext.EndGroups
}
