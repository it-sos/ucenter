package config

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"ucenter/s"
)

var ext = ".yml"

type Config struct {
	configName   string
	configFile   string
	configType   string
	envPrefix    string
	configPrefix string
}
// https://xuchao918.github.io/2019/04/29/%E4%BD%BF%E7%94%A8go%E8%AF%BB%E5%8F%96%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6/
// https://www.cnblogs.com/zhaof/p/8955332.html
func (c *Config) setYamlFilePathInfo() {
	cfg := strings.Split(c.configPrefix, ".")
	base := fmt.Sprintf("%s/%s/%s/%s", s.AppRoot, "config", c.envPrefix, s.AppEnv)
	for _, v := range cfg {
		base = filepath.Join(base, v)
		if _, err := os.Stat(base + ext); !os.IsNotExist(err) {
			break
		}
		if _, err := os.Stat(base); os.IsNotExist(err) {
			golog.Fatalf("config file not found, %s", base)
		}
	}

}

func (c *Config) initViper() {
	viper.SetConfigName(c.configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (c *Config) ConfServer(configPrefix string) {
	c.envPrefix = "server"
	c.configPrefix = configPrefix
	c.setYamlFilePathInfo()
	c.initViper()
}

//func (c *Config) ConfSecurity(file string, structs interface{}) interface{} {
//	c.pathDir = "security"
//	c.file = file
//	c.getYamlAbsPath()
//	v := c.parseYAML(structs)
//	fmt.Println(structs)
//	fmt.Println(v)
//	return v
//}
