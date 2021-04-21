package services

import (
	"testing"
	"ucenter/repositories"
	_ "ucenter/s/tests/testsimple"
)

func userConnect() UserService {
	return NewUserService(repositories.NewUserRepository())
}

func Test_userService_GetAll(t *testing.T) {
	t.Log(userConnect().GetByAccount("peng.yu"))
}
