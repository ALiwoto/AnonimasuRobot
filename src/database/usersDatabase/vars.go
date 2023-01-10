package usersDatabase

import (
	ws "github.com/AnimeKaizoku/ssg/ssg"
)

// database models
var (
	ModelBlockedUser = &BlockedUser{}
	BLockedUserEmpty = &BlockedUser{}
)

// caching
var (
	blockedUserMap   = ws.NewSafeEMap[int64, BlockedUser]()
	messageToUserMap = _getMessageToUserMap()
)
