package datamodels

import "time"

// 角色成员表
type UserRole struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`
	UserId     uint      `json:"user_id" xorm:"notnull unique 'uk_userid_roleid_appid' comment('用户ID')"` // 用户ID
	RoleId     uint      `json:"role_id" xorm:"notnull unique 'uk_userid_roleid_appid' comment('角色ID')"` //  角色ID
	AppId      uint      `json:"app_id" xorm:"notnull unique 'uk_userid_roleid_appid' comment('应用ID')"`  // 应用ID
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
