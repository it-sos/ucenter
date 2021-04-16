package sqlite

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"ucenter/s/db"
)

type sqlite struct{}

func (s *sqlite) GetDsn() string {
	return fmt.Sprintf("%s?loc=%s", db.Config.GetStorageFile(), db.Config.GetTimezone())
}

const driver = "sqlite3"

func (s *sqlite) Connect() (*xorm.EngineGroup, error) {
	return xorm.NewEngineGroup(driver, []string{s.GetDsn()})
}

func NewSqlite() db.Db {
	return &sqlite{}
}
