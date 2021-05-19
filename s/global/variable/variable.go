package variable

import (
	"os"
	"strings"
	"ucenter/s/global/consts"
)

var (
	// BasePath 项目根目录
	BasePath string = consts.TestBasePath
)

func init() {
	if dir, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
		} else {
			BasePath = dir
		}
	}

}
