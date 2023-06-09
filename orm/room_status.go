package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type RoomStatus struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name  string `db:"name" json:"name" gorm:"type:varchar(50)"`
	Image string `db:"image" json:"image" gorm:"type:varchar(100)"`
}
