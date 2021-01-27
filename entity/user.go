package entity

import (
	"time"
)

type User struct {
	ID         int64      `db:"id" json:"id"`
	Account    string     `db:"account" json:"account"`
	Password   string     `db:"password" json:"password"`
	Nickname   string     `db:"nickname" json:"nickname"`
	Phone      string     `db:"phone" json:"phone"`
	Expired    int64      `db:"expired" json:"expired"`
	IsDisabled uint8      `db:"is_disabled" json:"is_disabled"`
	IsDeleted  uint8      `db:"is_deleted" json:"is_deleted"`
	Salt       string     `db:"account" json:"account"`
	UpdateTime *time.Time `db:"update_time" json:"update_time"`
	CreateTime *time.Time `db:"create_time" json:"create_time"`
}
