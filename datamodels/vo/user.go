package vo

import (
	"ucenter/datamodels"
)

type User struct {
	datamodels.User
	DisabledName string `json:"disabledName" example:"已禁用"`
}

type UserVO struct {
	Paging
	Data []User `json:"data"` // 每页数据
}
