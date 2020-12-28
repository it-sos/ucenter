package config

import (
	"path/filepath"
	"testing"
	"ucenter/s"
)

func TestConfig_ConfSecurity(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
	}
	s.AppRoot = root
	s.AppEnv = "dev"
	//config := Config{}
	//c := config.ConfServer("mysql", Mysql{})
	//t.Log(c)
}
