package tests

import (
	"os"
	"ucenter/s/config"
)

func init() {
	os.Chdir("/data1/htdocs/ucenter")
	config.C.Init()
}
