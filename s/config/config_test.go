package config

import (
	"github.com/spf13/viper"
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
	config := Config{}
	config.ConfServer("mysql")
	t.Log(viper.Get("master.password"))
}
