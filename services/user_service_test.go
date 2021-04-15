package services

import (
	"testing"
	"ucenter/repositories"
	"ucenter/s/tests"
)

func userConnect() UserService {
	db := tests.ConnectDb()
	return NewUserService(repositories.NewUserRepository(db))
}

func Test_userService_GetAll(t *testing.T) {
	t.Log(userConnect().GetByAccount("peng.yu"))
}
