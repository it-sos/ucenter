package config

import (
	"fmt"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"ucenter/s"
)

var ext = ".yml"

type Config struct {
	pathDir     string
	yamlAbsPath string
	name        string
	structs     interface{}
	data        interface{}
}

// https://www.cnblogs.com/zhaof/p/8955332.html
func (c *Config) getYamlAbsPath() {
	cfg := strings.Split(c.name, ".")
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

func (c *Config) parseYAML() {
	data, err := ioutil.ReadFile(c.yamlAbsPath)
	if err != nil {
		golog.Fatal("parse yaml: %w", err)
	}

	if err := yaml.Unmarshal(data, &c.data); err != nil {
		golog.Fatal("parse yaml: %w", err)
	}
}

func (c *Config) parseStruct() interface{} {
	//structValue := reflect.Indirect(reflect.ValueOf(c.structs))
	//fmt.Println(structValue.Kind())
	v := c.data
	if reflect.TypeOf(v).String() == "map[string]interface {}" {
		data := v.(map[string]interface{})
		return data
	}
	//val := reflect.ValueOf(c.data)
	//val := reflect.ValueOf(c.structs)
	//vint := val.Interface()
	//fmt.Println(val.Kind())
	//fmt.Println(vint)
	//fmt.Println(structValue)
	//for key, value := range vint.(map[interface{}]interface{}) {
	//	fmt.Println(key, value)
	//}
	//fmt.Println(structValue)
	//fmt.Println(structValue.Elem())
	//structValue.Set(val)
}

func (c *Config) ConfServer(name string, structs interface{}) interface{} {
	c.pathDir = "server"
	c.name = name
	c.structs = structs
	c.getYamlAbsPath()
	c.parseYAML()
	c.parseStruct()
	return c.data
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
