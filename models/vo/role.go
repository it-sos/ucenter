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
