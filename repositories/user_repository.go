// file: repositories/movie_repository.go

package repositories

import (
	"sync"
	"ucenter/datamodels"
	"ucenter/s/db/common"
)

type UserRepository interface {
	// 新增
	Insert(p *datamodels.User) (id uint)
	// 更新
	Update(p *datamodels.User) (id uint)
	// 查询用户详细
	Select(p *datamodels.User) (datamodels.User, bool)
	// 查询用户列表
	SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User)
}

type userRepository struct {
	mu sync.RWMutex
	db *common.Db
}

var err error

func NewUserRepository(db *common.Db) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Select(p *datamodels.User) (datamodels.User, bool) {
	has, err := ur.db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (ur *userRepository) SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User) {
	user := make([]datamodels.User, 0)
	err := ur.db.Conn.Limit(limit, offset).Find(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func (ur *userRepository) Insert(p *datamodels.User) (id uint) {
	_, err = ur.db.Conn.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (ur *userRepository) Update(p *datamodels.User) (id uint) {
	_, err = ur.db.Conn.ID(p.Id).Update(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}
