package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"ucenter/s/db/mysql"
	"ucenter/s/db/sqlite"
)

import rds "ucenter/s/db/redis"

type Db interface {
	GetDsn() string
	Connect() *Dbs
}

type Dbs struct {
	Conn *xorm.EngineGroup
	Rdb  *redis.Client
}

var Conn *xorm.EngineGroup
var Rdb *redis.Client

func init() {
	if Config.GetStorageDriver() == driverMysql {
		Conn = mysql.NewMysql().Connect().Conn
	} else {
		Conn = sqlite.NewSqlite().Connect().Conn
	}
	Rdb = rds.NewRedis().Connect().Rdb
}
