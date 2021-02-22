package db

import (
	"github.com/go-xorm/xorm"
)

type Db struct {
	Conn *xorm.EngineGroup
}

type ConnectDb interface {
	GetDsn(dbType string) string
	Connect() (*Db, error)
}
