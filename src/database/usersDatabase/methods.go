package usersDatabase

import (
	"time"

	"github.com/AnimeKaizoku/AnonimasuRobot/src/core/wotoConfig"
)

//---------------------------------------------------------

func (b *BlockedUser) IsExpired(expirationDuration time.Duration) bool {
	if expirationDuration < 0 {
		// duration less than 0 means a block on user will never get expired.
		return false
	}

	return time.Since(b.UpdatedAt) > wotoConfig.GetExpirationDays()
}

func (b *BlockedUser) IsValid(expirationDuration time.Duration) bool {
	return b != BLockedUserEmpty && !b.IsExpired(expirationDuration)
}

// --------------------------------------------------------
