package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
)

type Db struct {
	Conn *xorm.EngineGroup
	Rdb  *redis.Client
}

type ConnectDb interface {
	GetDsn(dbType string) string
	Connect() (*Db, error)
}
