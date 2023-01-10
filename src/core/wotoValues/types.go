package wotoValues

import (
	"sync"

	"gorm.io/gorm"
)

type WotoCore struct {
	DbMutex *sync.Mutex
	DB      *gorm.DB
}

type UserPermission uint8

type (
	ButtonsUniqueId      string
	TextEntitiesUniqueId string
	TextEntitiesGroupId  string
)
