package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Message struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"` // use from firebase
	orm.GormModel

	RoomID        string `db:"room_id" json:"room_id" gorm:"type:varchar(36);index;"`
	UserID        string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
	Message       string `db:"message" json:"message" gorm:"type:text"`
	MessageTypeID string `db:"message_type_id" json:"message_type_id" gorm:"type:varchar(36)"`
}
