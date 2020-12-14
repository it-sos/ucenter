package config

import (
	"path/filepath"
	"testing"
	"ucenter/s"
)

type Mysql struct {
	Master map[string]string `yaml:"Master"`
	Slave1 map[string]string `yaml:"Slave1"`
	Slave2 map[string]string `yaml:"Slave2"`
	backup map[string]string `yaml:"backup"`
}

func TestConfig_ConfSecurity(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
	}
	s.AppRoot = root
	s.AppEnv = "dev"
	config := Config{}
	c := config.ConfServer("mysql", Mysql{})
	t.Log(c)
}
