package usersDatabase

import (
	ws "github.com/ALiwoto/ssg/ssg"
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
