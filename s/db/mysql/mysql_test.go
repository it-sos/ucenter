package mysql

import (
	"path/filepath"
	"testing"
	"ucenter/s"
)

type Role struct {
	Id         int64 `xorm:"pk id"` //指定主键并自增
	Name       string
	Info       string
	UpdateTime int64 `xorm:"update_time"` //修改后自动更新时间
	CreateTime int64 `xorm:"create_time"` //创建时间
}

func TestMysql(t *testing.T) {
	root, err := filepath.Abs("../../../")
	if err != nil {
	}
	s.AppRoot = root
	s.AppEnv = "dev"
	x, _ := ConnectMySQL()
	role := new(Role)
	role.Name = "管理员"
	role.Info = "超级管理员"
	affected, err := x.Conn.Insert(role)
	t.Log(affected, err)
}
