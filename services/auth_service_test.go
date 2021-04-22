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

func Test_authService_Login(t *testing.T) {
	instance := authInstance()
	auth := vo.AuthVO{Account: "peng.yu", Password: "1234"}
	if _, err := instance.Login(auth); err != nil && err.Error() != "{\"errCode\":4001002,\"message\":\"用户名或密码错误\"}" {
		t.Fail()
	}
}
