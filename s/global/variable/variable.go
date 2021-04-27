package variable

import (
	"log"
	"os"
	"strings"
	"ucenter/s/errors"
)

var (
	// 项目根目录
	BasePath string
)

func init() {
	if path, err := os.Getwd(); err == nil {
		// 兼容单元测试程序程序启动时不同的路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(errors.Error("get_base_path_err").Error())
	}
}
