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

func initConfig() {
	root, _ := filepath.Abs("../../../")
	s.AppRoot = root
	s.AppEnv = "dev"
}

// add
func TestAdd(t *testing.T) {
	initConfig()
	x, _ := ConnectMySQL()
	role := new(Role)
	role.Name = "普通"
	role.Info = "超级普通"
	a, e := x.Conn.Insert(role)
	t.Log(a, e)
	if a == 1 {
		b, err := x.Conn.ID(1).Update(role)
		t.Log(b, err)
	}
	//r, err := x.Conn.Query("select * from role")
	//if r == "1" {
	//	t.Log(results, err)
	//}
}

// delete
func TestDelelte(t *testing.T) {
	initConfig()
	x, _ := ConnectMySQL()
	r, err := x.Conn.ID(1).Delete(new(Role))
	t.Log(r, err)
}

// update
func TestUpdate(t *testing.T) {
	initConfig()
	x, _ := ConnectMySQL()
	role := new(Role)
	role.Name = "普通0"
	role.Info = "超级普通0"
	r, err := x.Conn.ID(1).Update(role)
	t.Log(r, err)
}

// select
func TestSelect(t *testing.T) {
	initConfig()
	x, _ := ConnectMySQL()
	role := new(Role)
	role.Name = "普通0"
	role.Info = "超级普通0"
	r, err := x.Conn.Table("role").ID(1).Get(role)
	t.Log(r, err)
	t.Log(role.Name)
}
