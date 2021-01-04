package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"ucenter/s"
)

type Config struct {
	configName   string
	configFile   string
	configType   string
	envPrefix    string
	configPrefix string
}

func (c *Config) configSplit() {
	cfg := strings.Split(c.configPrefix, ".")
	base := fmt.Sprintf("%s/%s/%s/%s", s.AppRoot, "config", c.envPrefix, s.AppEnv)
	size := len(cfg)
	c.configFile = fmt.Sprintf("%s/%s", base, strings.Join(cfg[0:size-1], "/"))
	c.configName = cfg[size-1]
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
	c.configSplit()
	c.initViper()
}

func (c *Config) ConfSecurity(configPrefix string) {
	c.envPrefix = "server"
	c.configPrefix = configPrefix
	c.configSplit()
	c.initViper()
}
