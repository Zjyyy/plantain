package sqlite

import (
	"plantain/base"

	"gorm.io/gorm"
)

func CreateMockData(db *gorm.DB) {
	AddDriverListItem(db, &base.PDriverInDatabase{})
}
