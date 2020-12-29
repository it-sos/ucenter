package mysql

import (
	"context"
	"database/sql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"ucenter/s/config"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#
type MySQL struct {
	Conn *xorm.EngineGroup
}

func getDbTypeConfig(dbType string) (string, string, string, string) {
	return viper.GetString(dbType + ".user"),
		viper.GetString(dbType + ".password"),
		viper.GetString(dbType + ".host"),
		viper.GetString(dbType + ".database")
}

func getConfig() (string, string, string) {
	c := config.Config{}
	c.ConfServer("mysql")
	user, password, host, database := getDbTypeConfig("master")
	masterDataSourceName := ""
	user, password, host, database := getDbTypeConfig("slave1")
	slave1DataSourceName := ""
	user, password, host, database := getDbTypeConfig("slave2")
	slave2DataSourceName := ""
	//return user, password, host, database
}

func ConnectMySQL(dsn string) (*MySQL, error) {
	dataSourceNameSlice := []string{masterDataSourceName, slave1DataSourceName, slave2DataSourceName}
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

func (db *MySQL) Query(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := db.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if scannable, ok := dest.(Scannable); ok {
		return scannable.Scan(rows)
	}

	if !rows.Next() {
		return ErrNoRows
	}
	return rows.Scan(dest)
}

func (db *MySQL) QueryOne() {
}

func (db *MySQL) QSql() {
}

func (db *MySQL) QSqlOne() {
}

func (db *MySQL) Count() {
}

func (db *MySQL) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.Conn.ExecContext(ctx, query, args...)
}

func (db *MySQL) Transaction() {
}

func (db *MySQL) Commit() {
}

func (db *MySQL) Rollback() {
}
