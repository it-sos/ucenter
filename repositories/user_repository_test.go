package repositories

import (
	"github.com/google/uuid"
	"testing"
	"ucenter/datamodels"
	"ucenter/s/tests"
)

var user = &datamodels.User{Account: "peng.yu"}

func connect() UserRepository {
	db := tests.ConnectDb()
	return NewUserRepository(db)
}

func Test_userRepository_InsertOrUpdate(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	user.Uuid = uuid.String()
	t.Log(connect().Insert(user))
}

func Test_userRepository_SelectMany(t *testing.T) {
	t.Log(connect().SelectMany(user, 0, 3))
}

func Test_userRepository_Select(t *testing.T) {
	t.Log(connect().Select(user))
}
