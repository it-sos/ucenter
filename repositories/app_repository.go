// file: repositories/movie_repository.go

package repositories

import (
	"sync"
	"ucenter/datamodels"
	"ucenter/s/db"
)

type AppRepository interface {
	Select(p *datamodels.App) (datamodels.App, bool)
	SelectMany(p *datamodels.App, offset int, limit int) (results []datamodels.App)

	InsertOrUpdate(p *datamodels.App) (id int64)
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

func NewAppRepository(db *db.Db) AppRepository {
	return &appRepository{db: db}
}

func (pu *appRepository) Select(p *datamodels.App) (datamodels.App, bool) {
	has, err := pu.db.Conn.Get(p)
	if err != nil {
		panic(err)
	}
	return *p, has
}

func (pu *appRepository) SelectMany(p *datamodels.App, offset int, limit int) (results []datamodels.App) {
	app := make([]datamodels.App, 0)
	err := pu.db.Conn.Limit(limit, offset).Find(&app)
	if err != nil {
		panic(err)
	}
	return app
}

func (pu *appRepository) InsertOrUpdate(p *datamodels.App) (id int64) {
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
