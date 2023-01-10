package usersDatabase

import (
	"time"

	"github.com/AnimeKaizoku/AnonimasuRobot/src/core/wotoConfig"
)

//---------------------------------------------------------

func (b *BlockedUser) IsExpired() bool {
	return time.Since(b.UpdatedAt) > wotoConfig.GetExpirationDays()
}

func (b *BlockedUser) IsValid() bool {
	return b != BLockedUserEmpty && !b.IsExpired()
}

// --------------------------------------------------------
