package orm

import (
	"time"

	orm "github.com/FourWD/middleware/orm"
)

type LoginBody struct {
	Username string `db:"username" json:"username" `
	Password string `db:"password" json:"password" `
}

type LogiReturn struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Username    string    `db:"username" json:"username" `
	Firstname   string    `db:"firstname" json:"firstname" `
	Lastname    string    `db:"lastname" json:"lastname" `
	Nickname    string    `db:"nickname" json:"nickname" `
	Birthday    time.Time `db:"birthday" json:"birthday" `
	Avatar      string    `db:"avatar" json:"avatar" `
	Email       string    `db:"email" json:"email" `
	Position    string    `db:"position" json:"position" `
	TokenExpire time.Time `json:"token_expire"`
}
