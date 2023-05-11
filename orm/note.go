package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Note struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"` // use from firebase
	orm.GormModel

	AppID  string `db:"app_id" json:"app_id" gorm:"type:varchar(36)"`
	RoomID string `db:"room_id" json:"room_id" gorm:"type:varchar(36);index;"`
	UserID string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
	Note   string `db:"note" json:"note" gorm:"type:text"`
}
