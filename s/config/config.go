package config

import (
	"fmt"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"ucenter/s"
)

type Mysql struct {
	Master map[string]string `yaml:"Master"`
	Slave1 map[string]string `yaml:"Slave1"`
	Slave2 map[string]string `yaml:"Slave2"`
	backup map[string]string `yaml:"backup"`
}

var ext = ".yml"

type Config struct {
	pathDir     string
	yamlAbsPath string
	file        string
}

func (c *Config) getYamlAbsPath() {
	cfg := strings.Split(c.file, ".")
	absFilepath := fmt.Sprintf("%s/%s/%s/%s", s.AppRoot, "config", c.pathDir, s.AppEnv)
	for _, v := range cfg {
		absFilepath = filepath.Join(absFilepath, v)
		if _, err := os.Stat(absFilepath + ext); !os.IsNotExist(err) {
			c.yamlAbsPath = absFilepath + ext
			break
		}
		if _, err := os.Stat(absFilepath); os.IsNotExist(err) {
			golog.Fatalf("config file not found, %s", absFilepath)
		}
	}
}

func (c *Config) parseYAML(structs interface{}) interface{} {
	data, err := ioutil.ReadFile(c.yamlAbsPath)
	if err != nil {
		golog.Fatal("parse yaml: %w", err)
	}

	if err := yaml.Unmarshal(data, &structs); err != nil {
		golog.Fatal("parse yaml: %w", err)
	}
	return structs
}

func (c *Config) ConfServer(file string, structs interface{}) interface{} {
	c.pathDir = "server"
	c.file = file
	c.getYamlAbsPath()
	return c.parseYAML(structs)
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
