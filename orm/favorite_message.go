package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type FavoriteMessage struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID  string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
	Title   string `db:"title" json:"title" gorm:"type:varchar(100)"`
	Message string `db:"message" json:"message" gorm:"type:text"`
}
