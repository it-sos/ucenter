package datamodels

import "time"

// 角色权限表
type RolePermission struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`
	RoleId     uint      `json:"role_id" xorm:"notnull unique 'uk_roleid_appid_routeid' comment('角色ID')"`  // 角色ID
	AppId      uint      `json:"app_id" xorm:"notnull unique 'uk_roleid_appid_routeid' comment('应用ID')"`   // 应用ID
	RouteId    uint      `json:"route_id" xorm:"notnull unique 'uk_roleid_appid_routeid' comment('路由ID')"` // 路由ID
	IsAllowed  uint8     `json:"is_allowed" xorm:"notnull default 1 comment('是否允许访问 0=否；1=是')"`            // 是否允许访问0=否；1=是
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
