package tests

import (
	"os"
	"ucenter/s/db"
	"ucenter/s/db/connect"
	"ucenter/web/bootstrap"
)

func ConnectDb() *db.Db {
	os.Chdir("/data1/htdocs/punch-in")
	bootstrap.SetupConfig()
	conn := connect.Connect()
	conn.Conn.ShowSQL(true)
	return conn
}
