package services

import (
	"testing"
	"ucenter/models/vo"
	"ucenter/repositories"
	_ "ucenter/s/tests/testsimple"
)

func authInstance() AuthService {
	return NewAuthService(NewUserService(repositories.NewUserRepository()))
}

func TestAuthService_LoginFail(t *testing.T) {
	i := authInstance()
	auth := vo.AuthVO{Account: "peng.yu", Password: "1234"}
	if _, err := i.Login(auth); err != nil && err.Error() != "{\"errCode\":4001002,\"message\":\"用户名或密码错误\"}" {
		t.Fail()
	}
}

func TestAuthService_Login(t *testing.T) {
	i := authInstance()
	auth := vo.AuthVO{
		Account:  "peng.yu",
		Password: "123456",
		Captcha:  "92554",
	}
	t.Log(i.Login(auth))
}

func TestAuthService_Register(t *testing.T) {
	i := authInstance()
	t.Log(i.Register(vo.UserParamsVO{
		Account:    "peng.yu",
		Password:   "123456",
		Nickname:   "大雨",
		Phone:      "13301350829",
		Expired:    0,
		IsDisabled: 0,
		IsDeleted:  0,
	}))
}

func TestAuthService_GenerateCaptcha(t *testing.T) {
	i := authInstance()
	t.Log(i.GenerateCaptcha("peng.yu"))
}
