package sqlite

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct{}

func (s *Sqlite) GetDsn(dbType string) string {
	return dbType + ".db?loc=Asia/Shanghai"
}

func (s *Sqlite) Connect() (*xorm.EngineGroup, error) {
	dataSourceNameSlice := []string{s.GetDsn("sqlite3")}
	return xorm.NewEngineGroup("sqlite3", dataSourceNameSlice)
}
