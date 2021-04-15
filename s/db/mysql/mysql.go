package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"ucenter/s/db"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#

type Mysql struct{}

func (m *Mysql) GetDsn(c db.Configuration) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%d?charset=%s&parseTime=true&loc=Local",
		c.GetUser(),
		c.GetPassword(),
		c.GetHost(),
		c.GetPort(),
		c.GetDatabase(),
		c.GetCharset(),
	)
}

func (m *Mysql) Connect() (*xorm.EngineGroup, error) {
	config := db.Configuration{}
	config.UseMysql()
	config.SetMode(db.MASTER)
	master := m.GetDsn(config)
	config.SetMode(db.SLAVE1)
	slave1 := m.GetDsn(config)
	config.SetMode(db.SLAVE2)
	slave2 := m.GetDsn(config)
	dataSourceNameSlice := []string{master, slave1, slave2}
	return xorm.NewEngineGroup("mysql", dataSourceNameSlice)
}
