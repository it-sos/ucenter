// file: services/movie_service.go

package services

import (
	"time"
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
}

//func NewUserService(repo repositories.UserRepository) UserService {
//	return &userService{
//		repo: repo,
//	}
//}

type userService struct {
	repo repositories.UserRepository
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

func (s *userService) GetByID(id uint) map[string]interface{} {
	user.Id = id
	return s.getUserInfo()
}

func (s *userService) GetByDate() map[string]interface{} {
	return s.getUserInfo()
}

/**
 * 获取当前时间，例：12:00
 */
func (s *userService) GetCurTime() string {
	var curTime string
	switch hour := time.Now().Hour(); {
	case hour < 10:
		curTime = "08:00"
	case hour < 12:
		curTime = "10:00"
	case hour < 14:
		curTime = "12:00"
	case hour < 16:
		curTime = "14:00"
	case hour < 18:
		curTime = "16:00"
	default:
		curTime = "18:00"
	}
	return curTime
}

/**
 * 获取当前星期数 例：Monday
 */
func (s *userService) GetCurWeek() string {
	return time.Now().Weekday().String()
}

func (s *userService) getUserInfo() map[string]interface{} {
	data, has := s.repo.Select(user)
	if false == has {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"id": data.Id,
	}
}

func (s *userService) UpdatePosterAndGenreByID(id uint, morning uint8, noon uint8, night uint8) uint {
	if id > 0 {
		user.Id = id
	}
	return 0
}
