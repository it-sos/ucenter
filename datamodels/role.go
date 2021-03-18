package datamodels

import "time"

// 角色表
type Role struct {
	Id         uint      `json:"id" readonly:"true" xorm:"notnull autoincr pk id"` // 角色id
	Name       string    `json:"name" xorm:"varchar(32) notnull comment('角色名称')"`  // 角色名称
	Info       string    `json:"info" xorm:"varchar(255) notnull comment('描述')"`   // 描述
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
