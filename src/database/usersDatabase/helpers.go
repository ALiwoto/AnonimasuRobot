package usersDatabase

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"time"

	"github.com/AnimeKaizoku/AnonimasuRobot/src/core/wotoConfig"
	wv "github.com/AnimeKaizoku/AnonimasuRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SetUserFromMessageId(messageId int64, user *gotgbot.User) {
	messageToUserMap.Add(messageId, user)
}

func GetUserFromMessageId(messageId int64) int64 {
	user := messageToUserMap.Get(messageId)
	if user == nil {
		return 0
	}

	return user.Id
}

func IsUserBlocked(userId int64) bool {
	expirationDuration := wotoConfig.GetExpirationDays()
	if expirationDuration == 0 {
		// 0 means block feature is disabled entirely.
		return false
	}

	if u := blockedUserMap.Get(userId); u != nil {
		return u.IsValid(expirationDuration)
	}

	encoded := _encodeUserId(userId)
	if wv.Core.DB == nil {
		return false
	}

	user := new(BlockedUser)
	err := wv.Core.DB.First(user, encoded).Error
	if err != nil || len(user.UserId) == 0 {
		blockedUserMap.Add(userId, BLockedUserEmpty)
		return false
	}

	blockedUserMap.Add(userId, user)
	return true
}

func BlockUser(userId int64) {
	if IsUserBlocked(userId) {
		return
	}

	user := &BlockedUser{
		UserId: _encodeUserId(userId),
	}

	wv.Core.LockDb()
	tx := wv.Core.DB.Begin()
	tx.Save(user)
	tx.Commit()
	wv.Core.UnlockDb()

	blockedUserMap.Add(userId, user)
}

func _encodeUserId(userId int64) string {
	b16 := []byte(ssg.ToBase16(userId))
	first := sha256.Sum256(b16)
	second := sha512.Sum512(b16)

	return _userEncodingPrefix +
		base64.StdEncoding.EncodeToString(first[:]) +
		base64.StdEncoding.EncodeToString(second[:]) + _userEncodingSuffix
}

func _getMessageToUserMap() *ssg.SafeEMap[int64, gotgbot.User] {
	m := ssg.NewSafeEMap[int64, gotgbot.User]()

	m.SetInterval(8 * time.Hour)
	m.SetExpiration(48 * time.Hour)

	return m
}
