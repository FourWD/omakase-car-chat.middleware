package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Notification struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	DomainID string `db:"domain_id" json:"domain_id" gorm:"type:varchar(36);"`

	FromUserID string `db:"from_user_id" json:"from_user_id" gorm:"type:varchar(36);"`
	ToUserID   string `db:"to_user_id" json:"to_user_id" gorm:"type:varchar(36);"`
	MessageID  string `db:"message_id" json:"message_id" gorm:"type:varchar(36);"`
	IsSent     bool   `db:"is_sent" json:"is_sent" gorm:"type:varchar(1)"`
	SentDate   string `db:"sent_date" json:"sent_date" gorm:"type:varchar(36)"`
}
