package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"ucenter/s/db"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#

type Mysql struct{}

func (m *Mysql) GetDsn(dbType string) string {
	dbType = "mysql." + dbType
	user, password, host, port, database, charset := viper.GetString(dbType+".user"),
		viper.GetString(dbType+".password"),
		viper.GetString(dbType+".host"),
		viper.GetString(dbType+".port"),
		viper.GetString(dbType+".database"),
		viper.GetString(dbType+".charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", user, password, host, port, database, charset)
}

func (m *Mysql) Connect() (*db.Db, error) {
	dataSourceNameSlice := []string{m.GetDsn("master"), m.GetDsn("slave1"), m.GetDsn("slave1")}
	engineGroup, err := xorm.NewEngineGroup("mysql", dataSourceNameSlice)

	if err != nil {
		return nil, err
	}
	return &db.Db{
		Conn: engineGroup,
	}, nil
}
