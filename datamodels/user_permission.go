package datamodels

import "time"

// 用户权限表
type UserPermission struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	UserId     uint      `xorm:"notnull unique 'uk_userid_appid_routeid' comment('用户表ID')"`
	AppId      uint      `xorm:"notnull unique 'uk_userid_appid_routeid' comment('应用表ID')"`
	RouteId    uint      `xorm:"notnull unique 'uk_userid_appid_routeid' comment('路由表ID')"`
	IsAllowed  uint8     `xorm:"notnull default 1 comment('是否允许访问 0=否；1=是')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
