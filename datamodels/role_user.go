package datamodels

import "time"

// 角色成员表
type RoleUser struct {
	Id         int64     `xorm:"notnull autoincr pk id"`
	RoleId     uint      `xorm:"notnull unique 'uk_roleid_userid' comment('角色表ID')"`
	UserId     uint      `xorm:"notnull unique 'uk_roleid_userid' comment('用户表ID')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
