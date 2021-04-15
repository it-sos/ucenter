package connect

import (
	"github.com/spf13/viper"
	"time"
	"ucenter/s/db"
	"ucenter/s/db/mysql"
	"ucenter/s/db/sqlite"
)

func Db() *db.Db {
	useDriver := viper.GetString("use-driver")

	if useDriver == "mysql" {
		connectDb = new(mysql.Mysql)
	} else {
		connectDb = new(sqlite.Sqlite)
	}

	db, err := connectDb.Connect()
	if err != nil {
		panic(err)
	}

	db.Conn.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	db.Conn.DatabaseTZ, _ = time.LoadLocation("Asia/Shanghai")
	return db
}
