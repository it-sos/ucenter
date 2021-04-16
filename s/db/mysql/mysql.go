package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"ucenter/s/db/common"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#

type mysql struct{}

func (m *mysql) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%d?charset=%s&parseTime=true&loc=Local",
		common.Config.GetUser(),
		common.Config.GetPassword(),
		common.Config.GetHost(),
		common.Config.GetPort(),
		common.Config.GetDatabase(),
		common.Config.GetCharset(),
	)
}

const driver = "mysql"

func (m *mysql) Connect() *common.Dbs {
	common.Config.UseMysql()
	common.Config.SetMode(common.Master)
	master := m.GetDsn()
	common.Config.SetMode(common.Slave1)
	slave1 := m.GetDsn()
	common.Config.SetMode(common.Slave2)
	slave2 := m.GetDsn()
	dataSourceNameSlice := []string{master, slave1, slave2}
	engine, err := xorm.NewEngineGroup(driver, dataSourceNameSlice)
	engine.TZLocation, _ = time.LoadLocation(common.Config.GetTimezone())
	engine.DatabaseTZ, _ = time.LoadLocation(common.Config.GetTimezone())
	if err != nil {
		panic(err)
	}
	return &common.Dbs{Conn: engine}
}

func NewMysql() common.Db {
	return &mysql{}
}
