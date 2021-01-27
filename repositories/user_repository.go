package repositories

import (
	"ucenter/entity"
)

type Query func(entity.User) bool

type UserRepository interface {
	Test(query Query)
}

type userMysqlRepository struct {
	source entity.User
}

func (m *userMysqlRepository) Test(query Query) {

}
