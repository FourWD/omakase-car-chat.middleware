package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Config struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Key      string `db:"key" json:"key" gorm:"type:varchar(50)"`
	KeyValue string `db:"key_value" json:"key_value" gorm:"type:varchar(200)"`
}
