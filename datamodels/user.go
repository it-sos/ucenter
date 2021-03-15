package datamodels

import "time"

// 用户表
type User struct {
	Id         uint      `json:"id" xorm:"notnull autoincr pk id"`
	Uuid       string    `json:"uuid" xorm:"varchar(64) notnull index comment('uuid')"`
	Account    string    `json:"account" xorm:"varchar(32) notnull index comment('账号')"`
	Password   string    `json:"password" xorm:"varchar(64) notnull"`
	Nickname   string    `json:"nickname" xorm:"varchar(16) notnull"`
	Phone      string    `json:"phone" xorm:"varchar(16) notnull default ''"`
	Expired    uint      `json:"expired" xorm:"notnull default 0 comment('有效期0=永久/unix时间戳')"`
	IsDisabled uint8     `json:"disabled" xorm:"notnull default 0 comment('是否禁用状态1=是;0=否')"`
	IsDeleted  uint8     `json:"deleted" xorm:"notnull default 0 comment('是否删除状态1=是;0=否')"`
	Salt       string    `json:"-" xorm:"varchar(32) notnull"`
	UpdateTime time.Time `json:"update_time" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" xorm:"notnull created"`
}
