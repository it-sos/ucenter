package caches

import (
	"fmt"
	"golang.org/x/net/context"
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
	var (
		str string
		err error
	)
	if str, err = db.Rdb.Get(context.Background(), a.k).Result(); err != nil {
		panic(err)
	}
	return str
}

func (a *authCaptchaCmd) Set(val string) {
	if err := db.Rdb.Set(context.Background(), a.k, val, ttlCaptcha).Err(); err != nil {
		panic(err)
	}
}

func (a *authCaptchaCmd) Clear() bool {
	err := db.Rdb.Del(context.Background(), a.k).Err()
	if err != nil {
		panic(err)
	}
	return true
}
