package datamodels

import "time"

// 角色表
type Role struct {
	Id         uint      `xorm:"notnull autoincr pk id"`
	Name       string    `xorm:"varchar(32) notnull comment('角色名称')"`
	Info       string    `xorm:"varchar(255) notnull comment('描述')"`
	UpdateTime time.Time `xorm:"notnull updated"`
	CreateTime time.Time `xorm:"notnull created"`
}
