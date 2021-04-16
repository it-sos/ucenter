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

func (r *Redis) GetDsn() string {
	return fmt.Sprintf("%s:%d", db.Config.GetHost(), db.Config.GetPort())
}

var ctx = context.Background()

func (r *Redis) Connect() *db.Dbs {
	db.Config.UseRedis()
	db.Config.SetMode(db.Master)

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.GetDsn(),
		Password: db.Config.GetPassword(),
		DB:       db.Config.GetDatabase(),
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	return &db.Dbs{Rdb: rdb}
}

func NewRedis() db.Db {
	return &Redis{}
}
