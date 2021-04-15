package caches

import "ucenter/s/db"

type AuthCache interface {
}

type authCache struct {
	rdb *db.Db
}
