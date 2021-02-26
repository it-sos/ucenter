// file: repositories/movie_repository.go

package repositories

import (
	"sync"
	"ucenter/datamodels"
	"ucenter/s/db"
)

type RoleRepository interface {
	Select(p *datamodels.Role) (datamodels.Role, bool)
	SelectMany(p *datamodels.Role, offset int, limit int) (results []datamodels.Role)

	InsertOrUpdate(p *datamodels.Role) (id int64)
	// 加入用户到角色
	JoinUser(userId int, roleId int) bool
	// 从角色中移除用户
	RemoveUser(userId int, roleId int) bool
	// 查询角色下成员列表
	SelectManyUser(roleId int) []datamodels.User
	// 给角色授权
	GrantPermission(roleId int, permissionId int) bool
}

type roleRepository struct {
	mu sync.RWMutex
	db *db.Db
}

func NewRoleRepository(db *db.Db) RoleRepository {
	return &roleRepository{db: db}
}

func (pu *roleRepository) Select(p *datamodels.Role) (datamodels.Role, bool) {
	has, err := pu.db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (pu *roleRepository) SelectMany(p *datamodels.Role, offset int, limit int) (results []datamodels.Role) {
	role := make([]datamodels.Role, 0)
	err := pu.db.Conn.Limit(limit, offset).Find(&role)
	if err != nil {
		panic(err)
	}
	return role
}

func (pu *roleRepository) InsertOrUpdate(p *datamodels.Role) (id int64) {
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
