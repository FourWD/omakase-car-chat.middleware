package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Notification struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	DomainID string `db:"domain_id" json:"domain_id" gorm:"type:varchar(36);"`

	NotiToken     string `db:"noti_token" json:"noti_token" gorm:"type:varchar(50);"`
	Message       string `db:"message" json:"message" gorm:"type:text"`
	MessageTypeID string `db:"message_type_id" json:"message_type_id" gorm:"type:varchar(36)"`
	IsSent        bool   `db:"is_sent" json:"is_sent" gorm:"type:varchar(1)"`
	SentDate      string `db:"sent_date" json:"sent_date" gorm:"type:varchar(36)"`
}
