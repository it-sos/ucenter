package config

import (
	"fmt"
	"github.com/spf13/viper"
	"ucenter/s/core"
	"ucenter/s/global/variable"
)

type config struct {
	cName string
	cType string
	cPath string
}

type Config interface {
	SetName(cName string)
	SetType(cType string)
	SetPath(cPath string)
	GetFile(cFile string) string
	GetName() string
	GetType() string
	GetPath() string
	Init()
}

func (c *config) SetName(cName string) {
	c.cName = cName
}

func (c *config) SetType(cType string) {
	c.cType = cType
}

func (c *config) SetPath(cPath string) {
	c.cPath = cPath
}

func (c *config) GetFile(cFile string) string {
	return fmt.Sprintf("%s%s/%s", c.GetPath(), core.GetEnviron(), cFile)
}

func (c *config) GetName() string {
	return c.cName
}

func (c *config) GetType() string {
	return c.cType
}

func (c *config) GetPath() string {
	return c.cPath
}

func (c *config) Init() {
	viper.SetConfigName(c.GetName())
	viper.SetConfigType(c.GetType())
	viper.AddConfigPath(c.GetPath() + core.GetEnviron())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func newConfig() Config {
	return &config{}
}

var C Config

func init() {
	C = newConfig()
	C.SetName("config")
	C.SetType("yaml")
	C.SetPath(variable.BasePath + "/config/")
}
