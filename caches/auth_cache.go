package caches

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
	"ucenter/s/db"
)

type AuthCache interface {
	Key(k string) AuthCmd
}

type AuthCmd interface {
	Incr() int64
	Decr() int64
	Clear() bool
	expire()
}

type authCache struct{}

type authCmd struct {
	k string
}

func NewAuthCache() AuthCache {
	return authCache{}
}

const prefix = "auth_times_%s"

var ctx = context.Background()

func (a authCache) Key(k string) AuthCmd {
	return authCmd{fmt.Sprintf(prefix, k)}
}

const ttl = 3 * time.Hour

func (a authCmd) expire() {
	if err := db.Rdb.Expire(ctx, a.k, ttl).Err(); err != nil {
		panic(err)
	}
}

func (a authCmd) Decr() int64 {
	decr, err := db.Rdb.Decr(ctx, a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return decr
}

func (a authCmd) Clear() bool {
	err := db.Rdb.Del(ctx, a.k).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func (a authCmd) Incr() int64 {
	incr, err := db.Rdb.Incr(ctx, a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return incr
}
