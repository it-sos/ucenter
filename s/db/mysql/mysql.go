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
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
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
	var dataSourceNameSlice []string
	common.Config.UseMysql()
	modes := []string{common.Master, common.Slave1, common.Slave2, common.Slave3}
	for _, mode := range modes {
		common.Config.SetMode(mode)
		if common.Config.GetHost() != "" {
			dataSourceNameSlice = append(dataSourceNameSlice, m.GetDsn())
		}
	}
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
