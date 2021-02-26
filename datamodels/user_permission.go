package datamodels

import "time"

// 用户权限表
type PermissionUser struct {
	Id           int64     `xorm:"notnull autoincr pk id"`
	UserId       uint      `xorm:"notnull unique 'uk_userid_permid' comment('用户表ID')"`
	PermissionId uint      `xorm:"notnull unique 'uk_userid_permid' comment('权限表ID')"`
	IsForbidden  uint8     `xorm:"notnull default 1 comment('是否禁止访问 0=否；1=是')"`
	UpdateTime   time.Time `xorm:"notnull updated"`
	CreateTime   time.Time `xorm:"notnull created"`
}
