package db

import "github.com/spf13/viper"

type Configuration struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"database"`
	Charset     string `yaml:"charset"`
	mode        string
	storageType string
}

const (
	HOST     = "host"
	PORT     = "port"
	USER     = "user"
	PASSWORD = "password"
	DATABASE = "database"
	CHARSET  = "charset"
)

func (c Configuration) parse() {
	c.Host = c.getFieldString(HOST)
	c.Port = c.getFieldInt(PORT)
	c.User = c.getFieldString(USER)
	c.Password = c.getFieldString(PASSWORD)
	c.Database = c.getFieldInt(DATABASE)
	c.Charset = c.getFieldString(CHARSET)
}

func (c Configuration) getFieldString(field string) string {
	return viper.GetString(c.getStorageType() + "." + c.getMode() + "." + field)
}

func (c Configuration) getFieldInt(field string) int {
	return viper.GetInt(c.getStorageType() + "." + c.getMode() + "." + field)
}

const (
	MASTER = "master"
	SLAVE1 = "slave1"
	SLAVE2 = "slave2"
	SLAVE3 = "slave3"
)

// 设置mode master/slave
func (c Configuration) SetMode(mode string) {
	c.mode = mode
}

const (
	REDIS  = "redis"
	SQLITE = "sqlite"
	MYSQL  = "mysql"
)

func (c Configuration) UseRedis() {
	c.storageType = REDIS
	c.parse()
}

func (c Configuration) UseSqlite() {
	c.storageType = SQLITE
	c.parse()
}

func (c Configuration) UseMysql() {
	c.storageType = MYSQL
	c.parse()
}

func (c Configuration) getMode() string {
	return c.mode
}

func (c Configuration) getStorageType() string {
	return c.storageType
}

func (c Configuration) GetHost() string {
	return c.Host
}

func (c Configuration) GetPort() int {
	return c.Port
}

func (c Configuration) GetPassword() string {
	return c.Password
}

func (c Configuration) GetDatabase() int {
	return c.Database
}

func (c Configuration) GetCharset() string {
	return c.Charset
}

func (c Configuration) GetUser() string {
	return c.User
}
