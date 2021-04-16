package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"ucenter/s/db/common"
	"ucenter/s/db/mysql"
	"ucenter/s/db/sqlite"
)

import rds "ucenter/s/db/redis"

var Conn *xorm.EngineGroup
var Rdb *redis.Client

func init() {
	if common.Config.GetStorageDriver() == common.DriverMysql {
		Conn = mysql.NewMysql().Connect().Conn
	} else {
		Conn = sqlite.NewSqlite().Connect().Conn
	}
	Rdb = rds.NewRedis().Connect().Rdb
}
