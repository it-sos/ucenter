package redis

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis

type Redis struct{}

var ctx = context.Background()

func (r *Redis) Connect() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.master.host") + ":" + viper.GetString("redis.master.port"),
		Password: viper.GetString("redis.master.password"),
		DB:       viper.GetInt("redis.master.database"),
	})
	err := rdb.Ping(ctx).Err()
	return rdb, err
}
