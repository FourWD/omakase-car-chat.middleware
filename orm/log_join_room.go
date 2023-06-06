package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type LogJoinRoom struct { // move to middleware-chat
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	UserID       string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
	RoomID       string `db:"room_id" json:"room_id" gorm:"type:varchar(36)"`
	ActivityType string `db:"activity_type" json:"activity_type" gorm:"type:varchar(3)"`
	//001 Add member
	//002 Online
	//003 Offline
}
