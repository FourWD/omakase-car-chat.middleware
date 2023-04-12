package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type BookingStatus struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BookingNO  string `db:"name" json:"name" gorm:"type:varchar(100)"`
	SellerID   string
	CustomerID string
}
