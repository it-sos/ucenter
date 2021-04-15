package services

import (
	"testing"
	"ucenter/repositories"
	"ucenter/s/tests"
)

func authInstance() AuthService {
	db := tests.ConnectDb()
	return NewAuthService(NewUserService(repositories.NewUserRepository(db)))
}

func Test_authService_Login(t *testing.T) {
	instance := authInstance()
	if instance.Login("peng.yu", "12334565").Error() != "{\"errCode\":4001002,\"message\":\"用户名或密码错误\"}" {
		t.Fail()
	}
}
