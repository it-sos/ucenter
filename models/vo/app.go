package vo

import (
	"ucenter/datamodels"
)

type AppVO struct {
	datamodels.App
}

type AppPageVO struct {
	PagingVO
	Data []AppVO `json:"data"` // 每页数据
}

type AppParamsVO struct {
	Name string `json:"name" binding:"required"` // 应用名称
	Info string `json:"info" binding:"required"` // 应用描述
	Link string `json:"link" binding:"required"` // 应用链接
	Icon string `json:"icon" binding:"required"` // 应用图标
}

type AppIconVO struct {
	Icon string `json:"icon" binding:"required"` // 应用图标标志
}
