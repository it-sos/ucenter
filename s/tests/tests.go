package tests

import (
	"os"
	"ucenter/web/bootstrap"
)

func init() {
	os.Chdir("/data1/htdocs/ucenter")
	bootstrap.SetupConfig()
	bootstrap.SetupInitDb()
}
