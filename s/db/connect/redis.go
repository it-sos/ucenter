package connect

import (
	"ucenter/s/db"
	"ucenter/s/db/redis"
)

func Redis() *db.Db {
	var redis db.ConnectDb = new(redis.Redis)
	db, err := redis.Connect()
	if err != nil {
		panic(err)
	}
	return db
}
