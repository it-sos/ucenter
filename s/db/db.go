package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"ucenter/s/db/mysql"
	"ucenter/s/db/sqlite"
)

type Db interface {
	GetDsn() string
	Connect() (*xorm.EngineGroup, error)
}

type dbs struct {
	conn *xorm.EngineGroup
	rdb  *redis.Client
}

type Dbs interface {
	SetConn(conn *xorm.EngineGroup)
	SetRdb(rdb *redis.Client)
}

var Conn *xorm.EngineGroup
var err error

func init() {
	if Config.GetStorageDriver() == driverMysql {
		Conn, err = mysql.NewMysql().Connect()
		if err != nil {
			panic(err)
		}
	} else {
		Conn, err = sqlite.NewSqlite().Connect()
		if err != nil {
			panic(err)
		}
	}
}
