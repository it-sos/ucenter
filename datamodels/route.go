package datamodels

import "time"

// 路由表
type Route struct {
	Id         uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`       // 路由ID
	AppId      uint      `json:"app_id" xorm:"notnull index comment('应用ID')"`                        // 应用ID
	Name       string    `json:"name" xorm:"varchar(32) notnull comment('功能名称')"`                    // 功能名称
	Info       string    `json:"info" xorm:"varchar(255) notnull comment('功能描述')"`                   // 功能描述
	Router     string    `json:"router" xorm:"varchar(255) notnull comment('路由')"`                   // 路由
	Method     string    `json:"method" xorm:"varchar(8) notnull comment('方法 PUT/DELETE/POST/GET')"` // 方法
	ParentId   uint      `json:"parent_id" xorm:"notnull default 0 comment('父ID')"`                  // 父ID
	UpdateTime time.Time `json:"update_time" readonly:"true" xorm:"notnull updated"`
	CreateTime time.Time `json:"create_time" readonly:"true" xorm:"notnull created"`
}
