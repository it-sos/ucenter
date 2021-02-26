package datamodels

import "time"

// 应用表
type App struct {
	Id         int64     `xorm:"notnull autoincr pk id"`
	Appid      string    `xorm:"varchar(32) notnull index comment('应用ID')"`
	Appsecret  string    `xorm:"varchar(32) notnull comment('应用秘钥')"`
	Name       string    `xorm:"varchar(32) notnull comment('应用名称')"`
	Info       string    `xorm:"varchar(255) notnull comment('应用描述')"`
	Link       string    `xorm:"varchar(255) notnull default '' comment('应用链接')"`
	Icon       string    `xorm:"varchar(255) notnull default '' comment('应用图标')"`
	IsDeleted  uint8     `xorm:"notnull default 0 comment('删除状态 1=是；0=否')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
