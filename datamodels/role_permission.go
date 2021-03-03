package datamodels

import "time"

// 角色权限表
type RolePermission struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	RoleId     uint      `xorm:"notnull unique 'uk_roleid_appid_routeid' comment('角色表ID')"`
	AppId      uint      `xorm:"notnull unique 'uk_roleid_appid_routeid' comment('应用表ID')"`
	RouteId    uint      `xorm:"notnull unique 'uk_roleid_appid_routeid' comment('路由表ID')"`
	IsAllowed  uint8     `xorm:"notnull default 1 comment('是否允许访问 0=否；1=是')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
