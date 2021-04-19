package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"sync"
	"ucenter/s/db/common"
	"ucenter/s/db/mysql"
	rds "ucenter/s/db/redis"
	"ucenter/s/db/sqlite"
)

var Conn *xorm.EngineGroup
var Rdb *redis.Client
var once sync.Once

func Init() {
	once.Do(func() {
		if common.Config.Init().GetStorageDriver() == common.DriverMysql {
			Conn = mysql.NewMysql().Connect().Conn
		} else {
			Conn = sqlite.NewSqlite().Connect().Conn
		}
		Rdb = rds.NewRedis().Connect().Rdb
	})
}
