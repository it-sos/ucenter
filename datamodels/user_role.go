package datamodels

import "time"

// 角色成员表
type UserRole struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	UserId     uint      `xorm:"notnull unique 'uk_userid_roleid_appid' comment('用户表ID')"`
	RoleId     uint      `xorm:"notnull unique 'uk_userid_roleid_appid' comment('角色表ID')"`
	AppId      uint      `xorm:"notnull unique 'uk_userid_roleid_appid' comment('应用表ID')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
