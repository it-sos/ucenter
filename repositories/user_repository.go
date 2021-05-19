// file: repositories/movie_repository.go

package repositories

import (
	"ucenter/datamodels"
	"ucenter/s/db"
)

type UserRepository interface {
	// Insert 新增
	Insert(p *datamodels.User) (id uint)
	// Update 更新
	Update(p *datamodels.User) (id uint)
	// Select 查询用户详细
	Select(p *datamodels.User) (datamodels.User, bool)
	// SelectMany 查询用户列表
	SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User)
}

type userRepository struct {
}

var err error

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Select(p *datamodels.User) (datamodels.User, bool) {
	has, err := db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (ur *userRepository) SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User) {
	user := make([]datamodels.User, 0)
	err := db.Conn.Limit(limit, offset).Find(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func (ur *userRepository) Insert(p *datamodels.User) (id uint) {
	_, err = db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *userRepository) Update(p *datamodels.User) (id uint) {
	_, err = db.Conn.ID(p.Id).Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}
