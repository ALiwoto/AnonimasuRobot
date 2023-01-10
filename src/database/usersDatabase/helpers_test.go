package usersDatabase_test

import (
	"testing"

	"github.com/AnimeKaizoku/AnonimasuRobot/src/database/usersDatabase"
)

func TestUserHistorySlices(t *testing.T) {
	usersDatabase.IsUserBlocked(10)
	usersDatabase.IsUserBlocked(10)
	usersDatabase.IsUserBlocked(10)
	usersDatabase.IsUserBlocked(20)
	usersDatabase.IsUserBlocked(30)
}
