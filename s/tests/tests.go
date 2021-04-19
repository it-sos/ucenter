package tests

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"ucenter/s/core"
)

func init() {
	os.Chdir("/data1/htdocs/ucenter")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/" + core.GetEnviron())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//bootstrap.SetupInitDb()
}
