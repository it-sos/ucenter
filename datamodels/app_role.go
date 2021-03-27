package datamodels

import "time"

// 应用角色表
type AppRole struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`
	AppId      uint      `json:"app_id" xorm:"notnull unique 'uk_appid_roleid' comment('应用ID')"`  // 应用ID
	RoleId     uint      `json:"role_id" xorm:"notnull unique 'uk_appid_roleid' comment('角色ID')"` // 角色ID
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
