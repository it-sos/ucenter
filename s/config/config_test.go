package config

import (
	"bytes"
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

func TestByte(t *testing.T) {
	s := []byte("")
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	t.Log(bytes.Equal(s1, s2))
}
