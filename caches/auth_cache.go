package caches

import (
	"ucenter/s/db/common"
)

type AuthCache interface {
}

type authCache struct {
	rdb *common.Db
}
