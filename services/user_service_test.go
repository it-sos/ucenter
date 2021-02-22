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
	t.Log(connect().GetAll(2, 3))
}

func Test_userService_GetById(t *testing.T) {
	t.Log(connect().GetByID(4))
}

func Test_userService_GetByDate(t *testing.T) {
	t.Log(connect().GetByDate())
}

func Test_userService_UpdatePosterAndGenreById(t *testing.T) {
	t.Log(connect().UpdatePosterAndGenreByID(4, 0, 1, 0))
}

func Test_userService_GetCurTime(t *testing.T) {
	t.Log(connect().GetCurTime())
}

func Test_userService_GetCurWeek(t *testing.T) {
	t.Log(connect().GetCurWeek())
}
