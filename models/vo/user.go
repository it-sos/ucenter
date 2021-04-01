package vo

import (
	"ucenter/datamodels"
)

type UserVO struct {
	datamodels.User
	DisabledName string `json:"disabled_name" example:"禁用"`       // 禁用状态
	ExpireDate   string `json:"expire_date" example:"2022-12-14"` // 用户有效期
}

type UserPageVO struct {
	PagingVO
	Data []UserVO `json:"data"` // 每页数据
}

type UserAppRolesVO struct {
	UserVO
	AppRole map[string][]string `json:"app_role"` // 用户应用角色
}

type UserDisabledVO struct {
	IsDisabled string `json:"is_disabled" binding:"required"` // 禁用标志0=否，1=是
}

type UserDeletedVO struct {
	IsDeleted string `json:"is_deleted" binding:"required"` // 删除标志0=否，1=是
}

type PasswordVO struct {
	Password string `json:"password" binding:"required"` // 密码
}

type UserParamsVO struct {
	Account    string `json:"account" binding:"required"` // 帐号
	Password   string `json:"password" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"` // 昵称
	Phone      string `json:"phone"`
	Expired    uint   `json:"expired" default:"0"`     // 有效期0=永久，unix时间戳
	IsDisabled uint8  `json:"is_disabled" default:"0"` // 禁用状态1=是，0=否
	IsDeleted  uint8  `json:"is_deleted" default:"0"`  // 删除标志1=是，0=否
}
