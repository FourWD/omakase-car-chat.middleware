package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type BookingStatus struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BookingNO  string `db:"booking_no" json:"booking_no" gorm:"type:varchar(50)"`
	SellerID   string `db:"seller_id" json:"seller_id" gorm:"type:varchar(36)"`
	CustomerID string `db:"customer_id" json:"customer_id" gorm:"type:varchar(36)"`
	RoomID     string `db:"room_id" json:"room_id" gorm:"type:varchar(36)"`
}
