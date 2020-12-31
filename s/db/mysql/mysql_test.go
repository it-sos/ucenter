package mysql

import (
	"path/filepath"
	"testing"
	"time"
	"ucenter/s"
)

type Role struct {
	Id         int64 `xorm:"pk id"` //指定主键并自增
	Name       string
	Info       string
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP <-"` //修改后自动更新时间
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP <-"` //创建时间
}

func TestMysql(t *testing.T) {
	root, err := filepath.Abs("../../../")
	if err != nil {
	}
	s.AppRoot = root
	s.AppEnv = "dev"
	x, _ := ConnectMySQL()
	role := new(Role)
	role.Name = "普通"
	role.Info = "超级普通"
	//a, e := x.Conn.Insert(role)
	//t.Log(a, e)
	//results, err := x.Conn.Query("select * from user")
	//x.Conn.Sync2(role)
	results, err := x.Conn.ID(1).Update(&role)
	t.Log(results, err)
}
