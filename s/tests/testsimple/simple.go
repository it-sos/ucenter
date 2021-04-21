package testsimple

import (
	"ucenter/s/db"
	_ "ucenter/s/tests"
)

func init() {
	db.Init()
}
