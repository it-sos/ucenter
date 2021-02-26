package datamodels

import "time"

// 用户表
type User struct {
	Id         uint64    `xorm:"notnull autoincr pk id"`
	Account    string    `xorm:"varchar(32) notnull index comment('账号')"`
	Password   string    `xorm:"varchar(64) notnull"`
	Nickname   string    `xorm:"varchar(16) notnull"`
	Phone      string    `xorm:"varchar(16) notnull default ''"`
	Expired    uint      `xorm:"notnull default 0 comment('有效期0=永久/unix时间戳')"`
	IsDisabled uint8     `xorm:"notnull default 0 comment('是否禁用状态1=是;0=否')"`
	IsDeleted  uint8     `xorm:"notnull default 0 comment('是否删除状态1=是;0=否')"`
	Salt       string    `xorm:"varchar(32) notnull"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
