package datamodels

import "time"

// 应用
type App struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`              // 应用ID
	Appid      string    `json:"appid" readonly:"true" xorm:"varchar(32) notnull index comment('应用AppID')"` // 应用appid
	Appsecret  string    `json:"-" readonly:"true" xorm:"varchar(32) notnull comment('应用秘钥')"`              // 应用密钥
	Name       string    `json:"name" xorm:"varchar(32) notnull comment('应用名称')"`                           // 应用名称
	Info       string    `json:"info" xorm:"varchar(255) notnull comment('应用描述')"`                          // 应用描述
	Link       string    `json:"link" xorm:"varchar(255) notnull default '' comment('应用链接')"`               // 应用链接
	Icon       string    `json:"icon" xorm:"varchar(255) notnull default '' comment('应用图标')"`               // 应用图标
	IsDeleted  uint8     `json:"deleted" xorm:"notnull default 0 comment('是否删除状态1=是;0=否')"`                 // 删除标志1=是，0=否
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
