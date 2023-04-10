package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type RoomMember struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	RoomID string `db:"room_id" json:"room_id" gorm:"type:varchar(36);index;"`
	UserID string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
}
