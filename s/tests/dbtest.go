package tests

import (
	"os"
	"ucenter/s/db"
	"ucenter/s/db/connect"
	"ucenter/web/bootstrap"
)

func ConnectDb() *db.Db {
	os.Chdir("/data1/htdocs/ucenter")
	bootstrap.SetupConfig()
	conn := connect.Db()
	conn.Conn.ShowSQL(true)
	return conn
}
