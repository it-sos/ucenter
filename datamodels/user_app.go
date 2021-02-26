package datamodels

import "time"

// 应用成员表
type AppUser struct {
	Id         int64     `xorm:"notnull autoincr pk id"`
	AppId      uint      `xorm:"notnull unique 'uk_appid_userid' comment('应用表ID')"`
	UserId     uint      `xorm:"notnull unique 'uk_appid_userid' comment('用户表ID')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
