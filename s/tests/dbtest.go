package tests

import (
	"os"
	"ucenter/s/db/common"
	"ucenter/s/db/connect"
	"ucenter/web/bootstrap"
)

func ConnectDb() *common.Db {
	os.Chdir("/data1/htdocs/ucenter")
	bootstrap.SetupConfig()
	conn := connect.Db()
	conn.Conn.ShowSQL(true)
	return conn
}
