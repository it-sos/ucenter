package datamodels

import "time"

// 用户权限表
type UserPermission struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`
	UserId     uint      `json:"user_id" xorm:"notnull unique(uar) comment('用户ID')"`            // 用户ID
	AppId      uint      `json:"app_id" xorm:"notnull unique(uar) comment('应用ID')"`             // 应用ID
	RouteId    uint      `json:"route_id" xorm:"notnull unique(uar) comment('路由ID')"`           // 路由ID
	IsAllowed  uint8     `json:"is_allowed" xorm:"notnull default 1 comment('是否允许访问 0=否；1=是')"` // 是否允许访问 0=否；1=是
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
