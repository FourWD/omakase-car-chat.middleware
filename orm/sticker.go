package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Sticker struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainID string `db:"domain_id" json:"domain_id" gorm:"type:varchar(36);"`

	StickerID string `db:"sticker_id" json:"sticker_id" gorm:"type:varchar(36)"`
	Sticker1  string `db:"sticker1" json:"sticker1" gorm:"type:varchar(50)"`
	Sticker2  string `db:"sticker2" json:"sticker2" gorm:"type:varchar(50)"`
	Sticker3  string `db:"sticker3" json:"sticker3" gorm:"type:varchar(50)"`
	Sticker4  string `db:"sticker4" json:"sticker4" gorm:"type:varchar(50)"`
	Sticker5  string `db:"sticker5" json:"sticker5" gorm:"type:varchar(50)"`
	Sticker6  string `db:"sticker6" json:"sticker6" gorm:"type:varchar(50)"`
	Sticker7  string `db:"sticker7" json:"sticker7" gorm:"type:varchar(50)"`
	Sticker8  string `db:"sticker8" json:"sticker8" gorm:"type:varchar(50)"`
}
