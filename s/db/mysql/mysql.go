package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"ucenter/s/db"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#

type mysql struct{}

func (m *mysql) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%d?charset=%s&parseTime=true&loc=Local",
		db.Config.GetUser(),
		db.Config.GetPassword(),
		db.Config.GetHost(),
		db.Config.GetPort(),
		db.Config.GetDatabase(),
		db.Config.GetCharset(),
	)
}

const driver = "mysql"

func (m *mysql) Connect() (*xorm.EngineGroup, error) {
	db.Config.UseMysql()
	db.Config.SetMode(db.Master)
	master := m.GetDsn()
	db.Config.SetMode(db.Slave1)
	slave1 := m.GetDsn()
	db.Config.SetMode(db.Slave2)
	slave2 := m.GetDsn()
	dataSourceNameSlice := []string{master, slave1, slave2}
	engine, err := xorm.NewEngineGroup(driver, dataSourceNameSlice)
	engine.TZLocation, _ = time.LoadLocation(db.Config.GetTimezone())
	engine.DatabaseTZ, _ = time.LoadLocation(db.Config.GetTimezone())
	return engine, err
}

func NewMysql() db.Db {
	return &mysql{}
}
