package vo

import (
	"ucenter/datamodels"
)

type UserVO struct {
	datamodels.User
	DisabledName string `json:"disabledName" example:"禁用"`       // 禁用状态
	ExpireDate   string `json:"expireDate" example:"2022-12-14"` // 用户有效期
}

type UserPageVO struct {
	PagingVO
	Data []UserVO `json:"data"` // 每页数据
}

type UserAppRolesVO struct {
	UserVO
	AppRole map[string][]string `json:"appRole"` // 用户应用角色
}

type UserDisabledVO struct {
	IsDisabled string `json:"disabled" binding:"required"` // 禁用标志0=否，1=是
}

type UserDeletedVO struct {
	IsDeleted string `json:"deleted" binding:"required"` // 删除标志0=否，1=是
}

type PasswordVO struct {
	Password string `json:"password" binding:"required"` // 密码
}

type UserParamsVO struct {
	Account    string `json:"account" binding:"required"` // 帐号
	Password   string `json:"password" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"` // 昵称
	Phone      string `json:"phone"`
	Expired    uint   `json:"expired" default:"0"`  // 有效期0=永久，unix时间戳
	IsDisabled uint8  `json:"disabled" default:"0"` // 禁用状态1=是，0=否
	IsDeleted  uint8  `json:"deleted" default:"0"`  // 删除标志1=是，0=否
}

// 关联角色
type UserAppRoleVO struct {
	UserId string `json:"user_id" binding:"required"` // 用户id
	RoleId string `json:"role_id" binding:"required"` // 角色id
	AppId  string `json:"app_id" binding:"required"`  // 应用id
}
