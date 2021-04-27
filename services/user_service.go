// file: services/movie_service.go

package services

import (
	"github.com/google/uuid"
	"ucenter/datamodels"
	"ucenter/models/vo"
	"ucenter/repositories"
)

var user = &datamodels.User{Account: "peng.yu"}

type UserService interface {
	New(paramsVO vo.UserParamsVO) (vo.UserVO, error)
	Update(id int, paramsVO vo.UserParamsVO) (vo.UserVO, error)
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

func (s *userService) New(paramsVO vo.UserParamsVO) (vo.UserVO, error) {
	passwordShal, salt := NewAuthService(s).generalPassword(paramsVO.Password)
	newUUID, _ := uuid.NewUUID()
	user := &datamodels.User{
		Uuid:       newUUID.String(),
		Account:    paramsVO.Account,
		Password:   passwordShal,
		Nickname:   paramsVO.Nickname,
		Phone:      paramsVO.Phone,
		Expired:    paramsVO.Expired,
		Salt:       salt,
		IsDisabled: paramsVO.IsDisabled,
		IsDeleted:  paramsVO.IsDeleted,
	}
	s.repo.Insert(user)

	DisabledName := "禁用"
	if paramsVO.IsDisabled == 0 {
		DisabledName = "启用"
	}
	return vo.UserVO{
		User:         *user,
		DisabledName: DisabledName,
		ExpireDate:   "",
	}, nil
}

func (s *userService) Update(id int, paramsVO vo.UserParamsVO) (vo.UserVO, error) {
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
