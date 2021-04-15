// file: services/movie_service.go

package services

import (
	"ucenter/datamodels"
	"ucenter/models/vo"
	"ucenter/repositories"
)

var user = &datamodels.User{Account: "peng.yu"}

type UserService interface {
	Save(id int, paramsVO vo.UserParamsVO) vo.UserVO
	Remove(id int) error
	SetDisabled(id int, disabledVO vo.UserDisabledVO) error
	SetPassword(id int, passwordVO vo.PasswordVO) error
	GetByIdOrUuid(id uint, uuid string) (vo.UserVO, error)
	GetList(page vo.PageVO) (vo.UserPageVO, error)
	GetUserAppRole(id uint) (vo.UserAppRolesVO, error)
	GetByAccount(account string) (datamodels.User, bool)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

func (s *userService) Save(id int, paramsVO vo.UserParamsVO) vo.UserVO {
	panic("implement me")
}

func (s *userService) Remove(id int) error {
	panic("implement me")
}

func (s *userService) SetDisabled(id int, disabledVO vo.UserDisabledVO) error {
	panic("implement me")
}

func (s *userService) SetPassword(id int, passwordVO vo.PasswordVO) error {
	panic("implement me")
}

func (s *userService) GetByIdOrUuid(id uint, uuid string) (vo.UserVO, error) {
	panic("implement me")
}

func (s *userService) GetList(page vo.PageVO) (vo.UserPageVO, error) {
	panic("implement me")
}

func (s *userService) GetUserAppRole(id uint) (vo.UserAppRolesVO, error) {
	panic("implement me")
}

func (s *userService) GetByAccount(account string) (datamodels.User, bool) {
	return s.repo.Select(&datamodels.User{Account: account})
}

func (s *userService) GetAll(page int, limit int) []datamodels.User {
	if page < 1 {
		page = 1
	}
	if page > 10000 {
		page = 10000
	}
	offset := (page - 1) * limit
	return s.repo.SelectMany(user, offset, limit)
}
