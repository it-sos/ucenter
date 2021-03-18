package vo

import (
	"ucenter/datamodels"
)

type UserVO struct {
	datamodels.User
	DisabledName string `json:"disabledName" example:"已禁用"`      // 禁用状态
	ExpireDate   string `json:"expireDate" example:"2022-12-14"` // 用户有效期
}

type UserPageVO struct {
	Paging
	Data []UserVO `json:"data"` // 每页数据
}
