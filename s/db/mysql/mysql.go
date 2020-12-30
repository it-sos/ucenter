package mysql

import (
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"ucenter/s/config"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#
type MySQL struct {
	Conn *xorm.EngineGroup
}

func getDsn(dbType string) string {
	user, password, host, port, database, charset := viper.GetString(dbType+".user"),
		viper.GetString(dbType+".password"),
		viper.GetString(dbType+".host"),
		viper.GetString(dbType+".port"),
		viper.GetString(dbType+".database"),
		viper.GetString(dbType+".charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, database, charset)
}

func ConnectMySQL() (*MySQL, error) {
	c := config.Config{}
	c.ConfServer("mysql")
	dataSourceNameSlice := []string{getDsn("master"), getDsn("slave1"), getDsn("slave1")}
	engineGroup, err := xorm.NewEngineGroup("mysql", dataSourceNameSlice)
	if err != nil {
		return nil, err
	}
	return &MySQL{
		Conn: engineGroup,
	}, nil
}

var ErrNoRows = sql.ErrNoRows

func (db *MySQL) Insert() {
}

func (db *MySQL) MultiInsert() {
}

func (db *MySQL) Delete() {
}

func (db *MySQL) Update() {
}

func (db *MySQL) Query() {
}

func (db *MySQL) QueryOne() {
}

func (db *MySQL) QSql() {
}

func (db *MySQL) QSqlOne() {
}

func (db *MySQL) Count() {
}

func (db *MySQL) Exec() {
}

func (db *MySQL) Transaction() {
}

func (db *MySQL) Commit() {
}

func (db *MySQL) Rollback() {
}
