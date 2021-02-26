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
	JoinUser(userId int, appId int) bool
	// 从角色中移除用户
	RemoveUser(userId int, appId int) bool
	// 查询角色下成员列表
	SelectManyUser(appId int) []datamodels.User
	// 给角色授权
	GrantPermission(appId int, permissionId int) bool
}

type appRepository struct {
	mu sync.RWMutex
	db *db.Db
}

func NewRoleRepository(db *db.Db) RoleRepository {
	return &appRepository{db: db}
}

func (pu *appRepository) Select(p *datamodels.Role) (datamodels.Role, bool) {
	has, err := pu.db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (pu *appRepository) SelectMany(p *datamodels.Role, offset int, limit int) (results []datamodels.Role) {
	app := make([]datamodels.Role, 0)
	err := pu.db.Conn.Limit(limit, offset).Find(&app)
	if err != nil {
		panic(err)
	}
	return app
}

func (pu *appRepository) InsertOrUpdate(p *datamodels.Role) (id int64) {
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
