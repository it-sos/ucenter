package datamodels

import "time"

// 角色权限表
type PermissionRole struct {
	Id           int64     `xorm:"notnull autoincr pk id"`
	RoleId       uint      `xorm:"notnull unique 'uk_roleid_permid' comment('角色表ID')"`
	PermissionId uint      `xorm:"notnull unique 'uk_roleid_permid' comment('权限表ID')"`
	IsForbidden  uint8     `xorm:"notnull default 1 comment('是否禁止访问 0=否；1=是')"`
	UpdateTime   time.Time `xorm:"notnull updated"`
	CreateTime   time.Time `xorm:"notnull created"`
}
