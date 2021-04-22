package caches

import (
	"fmt"
	"time"
	"ucenter/s/db"
)

type AuthTimes interface {
	Key(k string) AuthTimesCmd
}

type AuthTimesCmd interface {
	Get() int64
	Incr() int64
	Decr() int64
	Clear() bool
	expire()
}

type authTimes struct{}

type authTimesCmd struct {
	k string
}

func (a *authTimesCmd) Get() int64 {
	times, _ := db.Rdb.Get(ctx, a.k).Int64()
	return times
}

var NAuthTimes = NewAuthTimes()

func NewAuthTimes() AuthTimes {
	return &authTimes{}
}

const prefixTimes = "auth_times_%s"

func (a *authTimes) Key(k string) AuthTimesCmd {
	return &authTimesCmd{fmt.Sprintf(prefixTimes, k)}
}

const ttlTimes = 3 * time.Hour

func (a *authTimesCmd) expire() {
	if err := db.Rdb.Expire(ctx, a.k, ttlTimes).Err(); err != nil {
		panic(err)
	}
}

func (a *authTimesCmd) Decr() int64 {
	decr, err := db.Rdb.Decr(ctx, a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return decr
}

func (a *authTimesCmd) Clear() bool {
	err := db.Rdb.Del(ctx, a.k).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func (a *authTimesCmd) Incr() int64 {
	incr, err := db.Rdb.Incr(ctx, a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return incr
}
