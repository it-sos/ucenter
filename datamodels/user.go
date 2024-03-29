package datamodels

import "time"

// 用户表
type User struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`                              // 用户表ID
	Uuid       string    `json:"uuid" example:"5bbc-4ala-3dja-1djs-0aja" xorm:"varchar(64) notnull unique comment('uuid')"` // 用户uuid
	Account    string    `json:"account" xorm:"varchar(32) notnull unique comment('账号')"`                                   // 帐号
	Password   string    `json:"-" xorm:"varchar(64) notnull"`
	Nickname   string    `json:"nickname" xorm:"varchar(16) notnull"` // 昵称
	Phone      string    `json:"phone" xorm:"varchar(16) notnull default ''"`
	Expired    uint      `json:"expired" xorm:"notnull default 0 comment('有效期0=永久/unix时间戳')"`   // 有效期0=永久，unix时间戳
	IsDisabled uint8     `json:"is_disabled" xorm:"notnull default 0 comment('是否禁用状态1=是;0=否')"` // 禁用状态1=是，0=否
	IsDeleted  uint8     `json:"is_deleted" xorm:"notnull default 0 comment('是否删除状态1=是;0=否')"`  // 删除标志1=是，0=否
	Salt       string    `json:"-" xorm:"varchar(32) notnull"`                                  // 盐
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
