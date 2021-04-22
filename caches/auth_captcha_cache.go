package caches

import (
	"fmt"
	"time"
	"ucenter/s/db"
)

type AuthCaptcha interface {
	Key(k string) AuthCaptchaCmd
}

type AuthCaptchaCmd interface {
	Get() string
	Set(val string)
	Clear() bool
}

type authCaptcha struct{}

type authCaptchaCmd struct {
	k string
}

var NAuthCaptcha = NewAuthCaptcha()

func NewAuthCaptcha() AuthCaptcha {
	return &authCaptcha{}
}

const prefixCaptcha = "auth_captcha_%s"

func (a *authCaptcha) Key(k string) AuthCaptchaCmd {
	return &authCaptchaCmd{fmt.Sprintf(prefixCaptcha, k)}
}

const ttlCaptcha = 2 * time.Minute

func (a *authCaptchaCmd) Get() string {
	return db.Rdb.Get(ctx, a.k).String()
}

func (a *authCaptchaCmd) Set(val string) {
	db.Rdb.Set(ctx, a.k, val, ttlCaptcha)
}

func (a *authCaptchaCmd) Clear() bool {
	err := db.Rdb.Del(ctx, a.k).Err()
	if err != nil {
		panic(err)
	}
	return true
}
