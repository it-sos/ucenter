package sqlite

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"ucenter/s/db"
)

type Sqlite struct{}

func (s *Sqlite) GetDsn(dbType string) string {
	return dbType + ".db?loc=Asia/Shanghai"
}

func (s *Sqlite) Connect() (*db.Db, error) {

	dataSourceNameSlice := []string{s.GetDsn("sqlite3")}
	engineGroup, err := xorm.NewEngineGroup("sqlite3", dataSourceNameSlice)

	if err != nil {
		return nil, err
	}
	return &db.Db{
		Conn: engineGroup,
	}, nil
}
