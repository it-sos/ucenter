package sqlite

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"ucenter/s/db/common"
)

type sqlite struct{}

func (s *sqlite) GetDsn() string {
	return fmt.Sprintf("%s?loc=%s", common.Config.GetStorageFile(), common.Config.GetTimezone())
}

const driver = "sqlite3"

func (s *sqlite) Connect() *common.Dbs {
	engine, err := xorm.NewEngineGroup(driver, []string{s.GetDsn()})
	if err != nil {
		panic(err)
	}
	return &common.Dbs{Conn: engine}
}

func NewSqlite() common.Db {
	return &sqlite{}
}
