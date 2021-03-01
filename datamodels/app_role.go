package datamodels

import "time"

// 应用角色表
type AppRole struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	AppId      uint      `xorm:"notnull unique 'uk_appid_roleid' comment('应用表ID')"`
	RoleId     uint      `xorm:"notnull unique 'uk_appid_roleid' comment('角色表ID')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
