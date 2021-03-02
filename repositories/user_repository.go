// file: repositories/movie_repository.go

package repositories

import (
	"sync"
	"ucenter/datamodels"
	"ucenter/s/db"
)

type UserRepository interface {
	// 新增或更新
	InsertOrUpdate(p *datamodels.User) (id uint)
	// 查询用户详细
	Select(p *datamodels.User) (datamodels.User, bool)
	// 查询用户列表
	SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User)
	// 查询用户所属角色
	SelectRole(userId uint) []datamodels.Role
	// 查询用户具备的权限
	SelectPermission(userId uint) []datamodels.Permission
	// 给用户授权
	GrantPermission(userId uint, AppId uint, RouteId uint) bool
}

type punchRepository struct {
	mu sync.RWMutex
	db *db.Db
}

var err error

func NewUserRepository(db *db.Db) UserRepository {
	return &punchRepository{db: db}
}

func (pu *punchRepository) Select(p *datamodels.User) (datamodels.User, bool) {
	has, err := pu.db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (pu *punchRepository) SelectMany(p *datamodels.User, offset int, limit int) (results []datamodels.User) {
	punch := make([]datamodels.User, 0)
	err := pu.db.Conn.Limit(limit, offset).Find(&punch)
	if err != nil {
		panic(err)
	}
	return punch
}

func (pu *punchRepository) InsertOrUpdate(p *datamodels.User) (id int64) {
	if p.Id > 0 {
		_, err = pu.db.Conn.ID(p.Id).Update(p)
		if err != nil {
			panic(err)
		}
	} else {
		_, err = pu.db.Conn.Insert(p)
		if err != nil {
			panic(err)
		}
	}
	return p.Id
}
