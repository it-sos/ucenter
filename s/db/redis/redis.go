package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"ucenter/s/db"
)

// https://github.com/go-redis/redis

type Redis struct{}

var ctx = context.Background()

func (r *Redis) Connect() (*redis.Client, error) {
	db.Config.UseRedis()
	db.Config.SetMode(db.Master)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", db.Config.GetHost(), db.Config.GetPort()),
		Password: db.Config.GetPassword(),
		DB:       db.Config.GetDatabase(),
	})
	err := rdb.Ping(ctx).Err()
	return rdb, err
}
