package services

import (
	"testing"
	"ucenter/repositories"
	"ucenter/s/tests"
)

func connect() UserService {
	db := tests.ConnectDb()
	return NewUserService(repositories.NewUserRepository(db))
}

func Test_userService_GetAll(t *testing.T) {
	//t.Log(connect().GetAll(2, 3))
}
