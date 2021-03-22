package vo

import (
	"ucenter/datamodels"
)

type RoleVO struct {
	datamodels.Role
}

type RolePageVO struct {
	PagingVO
	Data []RoleVO `json:"data"` // 每页数据
}

type RoleParamsVO struct {
	Name string `json:"name" binding:"required"` // 角色名称
	Info string `json:"info" binding:"required"` // 角色描述
}

// 关联应用
type RoleAppVO struct {
	RoleId string `json:"role_id" binding:"required"` // 角色id
	AppId  string `json:"app_id" binding:"required"`  // 应用id
}

// 关联用户
type RoleAppUserVO struct {
	RoleId string `json:"role_id" binding:"required"` // 角色id
	AppId  string `json:"app_id" binding:"required"`  // 应用id
	UserId string `json:"user_id" binding:"required"` // 用户id
}
