package usersDatabase

import "time"

type BlockedUser struct {
	UserId    string    `json:"user_id" gorm:"primaryKey"`
	UpdatedAt time.Time `json:"updated_at"`
}
