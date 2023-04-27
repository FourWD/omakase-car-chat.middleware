package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Room struct {
	ID         string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	FirebaseID string `db:"firebase_id" json:"firebase_id" gorm:"type:varchar(25);"`
	orm.GormModel

	DomainID string `db:"domain_id" json:"domain_id" gorm:"type:varchar(36);"`
	Name     string `db:"name" json:"name" gorm:"type:varchar(100)"`
}
