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
	config := Config{}
	v := &Mysql{}
	c := config.ConfServer("a.b.c", v)
	mapInstance := make(map[string]interface{})
	mapInstance.Decode()
	t.Log(v.Master)
	t.Log(c)
}
