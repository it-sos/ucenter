package redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
	_ "ucenter/s/tests"
)

func initConfig() *redis.Client {
	db := NewRedis().Connect().Rdb
	return db
}

func TestInitConfig(t *testing.T) {
	initConfig()
}
