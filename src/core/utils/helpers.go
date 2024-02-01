package utils

import (
	"strings"
	"unicode/utf8"

	"github.com/ALiwoto/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func ExtractRepliedToMessageId(msg *gotgbot.Message) int64 {
	if !strings.Contains(msg.Text, "\n") {
		return 0
	}

	wholeLines := strings.SplitN(msg.Text, "\n", 2)
	linkLine := wholeLines[0]
	if !strings.HasPrefix(linkLine, "https://t.me/") ||
		!strings.Contains(linkLine, "/") {
		return 0
	}

	linkParts := strings.Split(linkLine, "/")
	result := ssg.ToInt64(linkParts[len(linkParts)-1])
	if result <= 0 {
		return 0
	}

	msg.Text = wholeLines[1]
	linkLen := int64(utf8.RuneCountInString(linkLine)) + 1
	newEntities := []gotgbot.MessageEntity{}
	for i := 0; i < len(msg.Entities); i++ {
		if msg.Entities[i].Offset < linkLen {
			continue
		}

		msg.Entities[i].Offset -= linkLen
		newEntities = append(newEntities, msg.Entities[i])
	}

	msg.Entities = newEntities
	return result
}
