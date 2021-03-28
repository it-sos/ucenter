package vo

import (
	"ucenter/datamodels"
)

type RouteVO struct {
	datamodels.Route
}

type RoutePageVO struct {
	PagingVO
	Data []RouteVO `json:"data"` // 每页数据
}

type RouteParamsVO struct {
	AppId    uint   `json:"app_id" binding:"required"` // 应用ID
	Name     string `json:"name" binding:"required"`   // 功能名称
	Info     string `json:"info" binding:"required"`   // 功能描述
	Router   string `json:"info" binding:"required"`   // 路由
	Method   string `json:"info" binding:"required"`   // 方法
	ParentId string `json:"info" binding:"required"`   // 父ID
}
