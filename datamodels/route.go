package datamodels

import "time"

// 路由表
type Route struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	AppId      uint      `xorm:"notnull index comment('应用表ID')"`
	Name       string    `xorm:"varchar(32) notnull comment('功能名称')"`
	Info       string    `xorm:"varchar(255) notnull comment('功能描述')"`
	Router     string    `xorm:"varchar(255) notnull comment('路由')"`
	Method     string    `xorm:"varchar(8) notnull comment('方法 PUT/DELETE/POST/GET')"`
	ParentId   uint      `xorm:"notnull default 0 comment('父ID')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
