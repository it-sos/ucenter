package vo

import (
	"ucenter/datamodels"
)

type UserPermissionVO struct {
	datamodels.UserPermission
}

type UserPermissionListVO struct {
	Data []UserPermissionVO `json:"data"` // 权限数据列表
}

type UserPermissionsVO struct {
	UserId    uint  `json:"user_id" binding:"required"`  // 用户ID
	AppId     uint  `json:"app_id" binding:"required"`   // 应用ID
	RouteId   uint  `json:"route_id" binding:"required"` // 路由ID
	IsAllowed uint8 `json:"is_allowed" default:"1"`      // 是否允许访问 0=否；1=是
}

type UserPermissionParamsVO struct {
	Params []UserPermissionsVO `json:"params"` // 关联权限，支持批量操作
}
